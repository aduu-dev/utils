package exe2

import "context"

// NewTestRunner creates a new Test Runner.
// It only stores commands inside its internal buffer.
func NewTestRunner() *TestRunner {
	return &TestRunner{}
}

// TestRunner is a runner which does not run anything,
// but just stores given commands.
type TestRunner struct {
	commands []Command
}

// Command is a test struct which stores a call.
// Options which are added by the Run-method itself are omitted.
type Command struct {
	Command []string
	Setting ExecuteSetting
}

// Commands returns all commands executed with the runner.
func (r *TestRunner) Commands() []Command {
	return r.commands
}

func (r *TestRunner) RunE(ctx context.Context, splitResult SplitResult, settings ...SettingsFunc) (err error) {
	setting := extractSettingsFromSlice(settings)

	r.commands = append(r.commands, Command{
		Command: splitResult.command(),
		Setting: setting,
	})

	return nil
}

func (r *TestRunner) RunWithOutputE(ctx context.Context, splitResult SplitResult, settings ...SettingsFunc) (string, error) {
	setting := extractSettingsFromSlice(settings)

	r.commands = append(r.commands, Command{
		Command: splitResult.command(),
		Setting: setting,
	})

	return "", nil
}

/*
func (r *TestRunner) Run(ctx context.Context, splitResult SplitResult, settings ...SettingsFunc) {
	r.commands = append(r.commands, splitResult.command())
}

func (r *TestRunner) RunWithOutputE(ctx context.Context, splitResult SplitResult, settings ...SettingsFunc) (out string, err error) {
	r.commands = append(r.commands, splitResult.command())

	return "", nil
}
*/

/*
func (r *TestRunner) RunWithOutputCombined(ctx context.Context, splitResult SplitResult, settings ...SettingsFunc) string {
	r.commands = append(r.commands, splitResult.command())

	return ""
}

func (r *TestRunner) RunWithOutputCombinedE(ctx context.Context, splitResult SplitResult, settings ...SettingsFunc) (out string, err error) {
	r.commands = append(r.commands, splitResult.command())

	return "", nil
}
*/
