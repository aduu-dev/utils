package exe

import "time"

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
