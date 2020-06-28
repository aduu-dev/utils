package exe2

import (
	"os"
	"os/exec"
	"time"
)

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

func extractSettingsFromSlice(settings []SettingsFunc) ExecuteSetting {
	setting := defaultSettings()

	for _, settingsFunc := range settings {
		settingsFunc(&setting)
	}

	return setting
}

func applySettings(cmd *exec.Cmd, setting ExecuteSetting) {
	if setting.dir != "" {
		cmd.Dir = os.ExpandEnv(setting.dir)
	}

	if setting.output {
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}
}

func defaultSettings() ExecuteSetting {
	return ExecuteSetting{}
}
