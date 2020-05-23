// Package exe enables template completion before command execution.
package exe // import "aduu.dev/utils/exe"

import (
	"context"
	"os/exec"
	"time"
)

// Run runs the function on the default Runner.
func Run(cmd *exec.Cmd, name string, settings ...SettingsFunc) {
	r := DefaultRunner()
	r.Run(cmd, name, settings...)
}

// RunE runs the function on the default Runner.
func RunE(cmd *exec.Cmd, name string, settings ...SettingsFunc) (err error) {
	r := DefaultRunner()
	return r.RunE(cmd, name, settings...)
}

// RunWithOutput runs the function on the default Runner.
func RunWithOutput(cmd *exec.Cmd, name string, settings ...SettingsFunc) string {
	r := DefaultRunner()
	return r.RunWithOutput(cmd, name, settings...)
}

// RunWithOutputE runs the function on the default Runner.
func RunWithOutputE(cmd *exec.Cmd, name string, settings ...SettingsFunc) (out string, err error) {
	r := DefaultRunner()
	return r.RunWithOutputE(cmd, name, settings...)
}

func search(text []rune, what string) int {
	whatRunes := []rune(what)

	for i := range text {
		found := true
		for j := range whatRunes {
			if text[i+j] != whatRunes[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}

// SplitCommand runs the function on the default Runner.
func SplitCommand(s string) []string {
	r := DefaultRunner()
	return r.SplitCommand(s)
}

/*
func splitCommand2(s string) []string {
	input := s
	var values []string

	for len(s) != 0 {
		if s[0] == '"' {
			next := strings.Index(s[1:], "\"")
			if next == -1 {
				panic("no matching \": " + input)
			}

			values = append(values, s[1:next+1])
			s = s[next+2:]
		} else {
			next := strings.Index(s, " ")
			if next == -1 {
				values = append(values, s)
				s = ""
				break
			}

			values = append(values, s[:next])
			s = s[next+1:]
		}
	}

	return values
}
*/

// Execute runs the function on the default Runner.
func Execute(s string, obj interface{}, settings ...SettingsFunc) *exec.Cmd {
	r := DefaultRunner()
	return r.Execute(s, obj, settings...)
}

// ExecuteE runs the function on the default Runner.
func ExecuteE(s string, obj interface{}, settings ...SettingsFunc) (*exec.Cmd, error) {
	r := DefaultRunner()
	return r.ExecuteE(s, obj, settings...)
}

// Execute runs the function on the default Runner.
func ExecuteWithContext(ctx context.Context, s string, obj interface{}, settings ...SettingsFunc) *exec.Cmd {
	r := DefaultRunner()
	return r.ExecuteWithContext(ctx, s, obj, settings...)
}

// ExecuteE runs the function on the default Runner.
func ExecuteEWithContext(ctx context.Context, s string, obj interface{}, settings ...SettingsFunc) (*exec.Cmd, error) {
	r := DefaultRunner()
	return r.ExecuteEWithContext(ctx, s, obj, settings...)
}

// ExecuteWithOutput runs the function on the default Runner.
func ExecuteWithOutput(s string, obj interface{}, settings ...SettingsFunc) string {
	r := DefaultRunner()
	return r.ExecuteWithOutput(s, obj, settings...)
}

// ExecuteWithOutputE runs the command on the default runner.
func ExecuteWithOutputE(s string, obj interface{}, settings ...SettingsFunc) (out string, err error) {
	r := DefaultRunner()
	return r.ExecuteWithOutputE(s, obj, settings...)
}

// SettingsFunc is a function which modifies the execution setting.
type SettingsFunc func(s *ExecuteSetting)

// ExecuteSetting holds options to apply to a command runner.
type ExecuteSetting struct {
	dir     string
	start   bool
	output  bool
	timeout time.Duration
	//executeAstemplate bool
}

// WithStart sets the command to run async.
func WithStart(s *ExecuteSetting) {
	s.start = true
}

// WithDir sets working directory to use when running the command.
func WithDir(path string) SettingsFunc {
	return func(s *ExecuteSetting) {
		s.dir = path
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

/*
// WithoutTemplateExecution does
func WithoutTemplateExecution(s *ExecuteSetting) {
	s.executeAstemplate = false
}
*/
