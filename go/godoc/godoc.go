// Package godoc includes a method to serve a given package path on a given port on localhost.
package godoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"golang.org/x/xerrors"

	"aduu.dev/utils/exe"
	"aduu.dev/utils/go/gomod"
	"aduu.dev/utils/projectpath"
	"aduu.dev/utils/tempdir"

	"github.com/skratchdot/open-golang/open"
)

// Server starts a godoc server showing the given package path.
type Server interface {
	Serve(packagePath string, port int) (err error)
}

// BaseSettings holds settings for the godoc server.
type Settings struct {
	openBrowser bool
}

// SettingsFunc is a function mutating attributes inside a given BaseSettings-structure.
type SettingsFunc func(s *Settings)

// WithOpenBrowser sets whether a browser should be opened.
func WithOpenBrowser(yesNo bool) SettingsFunc {
	return func(s *Settings) {
		s.openBrowser = yesNo
	}
}

// DefaultServer returns a godoc.Server with default settings.
func DefaultServer(settings ...SettingsFunc) Server {
	newSettings := &Settings{
		openBrowser: false,
	}

	for _, mutator := range settings {
		mutator(newSettings)
	}

	return &server{
		settings: *newSettings,
	}
}

type server struct {
	// TODO: Use sync channel here.
	stopped bool
	// Locks stopped so only one stops and shuts down. Without it I can lose all mounted data.
	mux sync.Mutex

	settings Settings

	cmd            *exec.Cmd
	mountDirectory string
}

// Run shows aduu.dev via godoc by running a godoc server on the given port.
func (s *server) Serve(packagePath string, port int) (err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("starting godoc failed: %w", err)
		}
	}()

	dir, err := tempdir.MakeTempDirE(gomod.DirectCallerFunctionNameWithPackageName())
	if err != nil {
		return
	}
	defer func() {
		err = tempdir.DeleteTempDirExistingError(dir, &err)
	}()

	if err = s.nfsSetup(); err != nil {
		return err
	}

	s.mountDirectory = filepath.Join(dir, "src", "aduu.dev")
	if err = os.MkdirAll(s.mountDirectory, 0755); err != nil {
		return err
	}

	ppath := projectpath.ProjectPath()
	ppath, err = filepath.EvalSymlinks(ppath)
	if err != nil {
		return err
	}
	aduuPath := filepath.Join(ppath, "aduu.dev")

	settings := struct {
		MountPath      string
		MountDirectory string
	}{
		MountPath:      aduuPath,
		MountDirectory: s.mountDirectory,
	}
	exe.Execute(`mount localhost:{{.MountPath}} {{.MountDirectory}}`, settings)

	if err = os.Setenv("GOPATH", dir); err != nil {
		return err
	}

	s.cmd = exe.Execute(`godoc -http=:{{.}}`, port, exe.WithStart)

	// Sleep before opening browser. Gives godoc time to index most packages (probably).
	time.Sleep(time.Second * 5)

	if s.settings.openBrowser {
		// Open godoc in browser.
		address := fmt.Sprintf("http://localhost:%d/pkg/%s", port, packagePath)
		//if err = browser.OpenURL(address); err != nil {
		if err = open.Run(address); err != nil {
			fmt.Printf("failed to open browser url (%s): %v\n", address, err)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, os.Interrupt)
	go func() {
		for sig := range c {
			switch sig {
			case syscall.SIGTERM:
				fmt.Println("SIGTERM received")

				err2 := s.syncedShutdown()
				// Do not override old error if another go routine did the shutdown and set the error already.
				if err2 != nil {
					err = err2
				}

				fmt.Println("Exiting early with err =", err)
				os.Exit(1)
			case os.Interrupt:
				err2 := s.syncedShutdown()
				// Do not override old error if another go routine did the shutdown and set the error already.
				if err2 != nil {
					err = err2
				}
			default:
				log.Println("Unknown signal:", sig)
			}
		}
	}()

	for !s.stopped {
		time.Sleep(time.Second)
	}

	return nil
}

func (s *server) nfsSetup() (err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("nfs setup failed: %w", err)
		}
	}()

	// Setup /etc/exports.
	filename := "/etc/exports"
	wantedLine := os.ExpandEnv(`/ -alldirs -mapall=$USER localhost`)

	// Load file and check if it contains our string already.
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	if !strings.Contains(string(content), wantedLine) {
		cmd := `sudo tee -a /etc/exports <<< "/ -alldirs -mapall=$USER localhost"`
		exe.Execute("bash -c '{{.}}'", cmd)

		// Insert line.
		var f *os.File
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		// In case WriteString fails.
		defer f.Close()

		if _, err = f.WriteString("/ -alldirs -mapall=$USER localhost"); err != nil {
			return err
		}
		if err = f.Close(); err != nil {
			return
		}
	}

	// Start nfs service.
	exe.Execute(`sudo launchctl start com.apple.rpcbind`, "")
	exe.Execute(`sudo nfsd start`, "")

	return nil
}

func (s *server) syncedShutdown() (err error) {
	s.mux.Lock()
	if !s.stopped {
		s.stopped = true
		err = s.shutdown()
	}
	s.mux.Unlock()
	return
}

func (s *server) shutdown() (err error) {
	fmt.Println("Shutting down s..")

	defer func() {
		if err != nil {
			err = xerrors.Errorf("shutdown failed: %w", err)
		}
	}()

	if err = s.cmd.Process.Kill(); err != nil {
		fmt.Printf("killing godoc failed: %v\n", err)
	}

	if _, err2 := exe.ExecuteE(`umount {{.}}`, s.mountDirectory); err2 != nil {
		if err != nil {
			err = xerrors.Errorf("failed to unmount (%v) after killing godoc process failed: %w", err2, err)
		} else {
			err = xerrors.Errorf("failed to unmount %s: %w", s.mountDirectory, err2)
		}
	}
	return
}
