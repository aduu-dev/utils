// Package shell offers a shell to execute code in.
package shell

import (
	"os"
	"os/exec"
)

// Shell contains methods to run Shell scripts.
//
// If an error occured then any later runs do not get executed.
type Shell interface {
	// Err returns the last error.
	Err() error

	// RunScript runs the given scrip in a shell.
	RunScript(string)
}

// Err returns the last error.
func (sh *shShell) Err() error {
	return sh.err
}

// New creates a new /bin/sh shell.
func New() Shell {
	return &shShell{}
}

type shShell struct {
	err error
}

func (sh *shShell) RunScript(script string) {
	if sh.err != nil {
		return
	}

	c := exec.Command("/bin/sh", "-c", script)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		sh.err = err
		return
	}
}
