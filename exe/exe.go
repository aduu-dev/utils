// Package exe enables template completion before command execution.
package exe // import "aduu.dev/utils/exe"

import (
	"context"
	"os/exec"
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
