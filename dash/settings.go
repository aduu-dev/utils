package dash

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"aduu.dev/utils/errors2"
	"k8s.io/klog/v2"
)

// SettingsFunc is a function which modifies the execution setting.
type SettingsFunc func(s *ExecuteSetting)

// ExecuteSetting holds options to apply to a command runner.
type ExecuteSetting struct {
	Dir     string
	start   bool
	timeout time.Duration
	//executeAstemplate bool

	// stdin/stdout:
	output     bool
	StdinFile  string
	StdoutFile string
	StderrFile string

	openFiles []*os.File
}

// WithStart sets the command to run async.
func WithStart(s *ExecuteSetting) {
	s.start = true
}

// WithDir sets working directory to use when running the command.
func WithDir(path string) SettingsFunc {
	return func(s *ExecuteSetting) {
		s.Dir = path
	}
}

// withOutput makes the command run with output.
func withOutput(s *ExecuteSetting) {
	s.output = true
}

// WithTimeout sets a maximum duration to wait.
func WithTimeout(duration time.Duration) SettingsFunc {
	return func(s *ExecuteSetting) {
		s.timeout = duration
	}
}

// WithStdinFile opens the given file for reading with Stdin.
func WithStdinFile(path string) SettingsFunc {
	return func(s *ExecuteSetting) {
		s.StdinFile = path
	}
}

// WithStdoutFile opens the given file for writing with Stdout.
// Does not append.
func WithStdoutFile(path string) SettingsFunc {
	return func(s *ExecuteSetting) {
		s.StdoutFile = path
	}
}

// WithStderrFile opens the given file for writing with Stderr.
// Does not append.
func WithStderrFile(path string) SettingsFunc {
	return func(s *ExecuteSetting) {
		s.StdoutFile = path
	}
}

/*
// WithoutTemplateExecution does
func WithoutTemplateExecution(s *ExecuteSetting) {
	s.executeAstemplate = false
}
*/

func extractSettingsFromSlice(settings []SettingsFunc) ExecuteSetting {
	setting := defaultSettings()

	for _, settingsFunc := range settings {
		settingsFunc(&setting)
	}

	return setting
}

var (
	errOutputAndStdoutFileOptsIncompatible = fmt.Errorf("err: setting stdout and returning output is incompatible")

	errStartAndOutputIncompatible = fmt.Errorf("err: the flags start and output are incompatible")
)

func createCommand(ctx context.Context,
	splitResult *SplitResult,
	setting *ExecuteSetting,
) (cmd *exec.Cmd, cancel func(), err error) {
	defer func() {
		if err != nil {
			// Closing again because we encountered errors.
			for _, reader := range setting.openFiles {
				if closeErr := reader.Close(); closeErr != nil {
					err = errors2.CombineErrors(err, closeErr)
				}
			}

			err = fmt.Errorf("failed to apply settings: %w", err)
		}
	}()

	cancel = func() {}
	if setting.timeout != 0 {
		ctx, cancel = context.WithTimeout(ctx, setting.timeout)
	}

	cmd = exec.CommandContext(ctx, splitResult.Name, splitResult.Args...)

	if setting.start && setting.output {
		return nil, cancel, errStartAndOutputIncompatible
	}

	if setting.Dir != "" {
		cmd.Dir = os.ExpandEnv(setting.Dir)
	}

	if setting.output && len(setting.StdoutFile) != 0 {
		return nil, cancel, errOutputAndStdoutFileOptsIncompatible
	}

	if len(setting.StdinFile) == 0 {
		klog.V(5).InfoS("Setting stdin to os.Stdin")
		cmd.Stdin = os.Stdin
	} else {
		klog.V(5).InfoS("Seting stdin to file",
			"file", setting.StdinFile,
		)

		stdin, err := os.OpenFile(setting.StdinFile, os.O_RDWR, 0755)

		if err != nil {
			klog.ErrorS(err, "Error opening stdin",
				"file", setting.StdinFile)
			return nil, cancel, err
		}

		cmd.Stdin = stdin
		klog.V(5).InfoS("Adding stdin to open files")
		setting.openFiles = append(setting.openFiles, stdin)
	}

	if !setting.output {
		if len(setting.StdoutFile) == 0 {
			klog.V(5).InfoS("Setting stdout to os.Stdout")

			cmd.Stdout = os.Stdout
		} else {
			klog.V(5).InfoS("Seting stdout to file",
				"file", setting.StdoutFile,
			)

			stdout, err := os.OpenFile(setting.StdoutFile, os.O_CREATE|os.O_WRONLY, 0755)

			if err != nil {
				klog.ErrorS(err, "Error opening stdout",
					"file", setting.StdoutFile)
				return nil, cancel, err
			}

			cmd.Stdout = stdout

			setting.openFiles = append(setting.openFiles, stdout)
		}
	}

	if len(setting.StderrFile) == 0 {
		klog.V(5).InfoS("Setting stderr to os.Stderr")

		cmd.Stderr = os.Stderr
	} else {
		klog.V(5).InfoS("Seting stdout to file",
			"file", setting.StdoutFile,
		)

		stderr, err := os.OpenFile(setting.StderrFile, os.O_CREATE|os.O_WRONLY, 0755)

		if err != nil {
			klog.ErrorS(err, "Error opening stderr",
				"file", setting.StderrFile)
			return nil, cancel, err
		}

		cmd.Stderr = stderr

		setting.openFiles = append(setting.openFiles, stderr)
	}

	return cmd, cancel, nil
}

func defaultSettings() ExecuteSetting {
	return ExecuteSetting{}
}
