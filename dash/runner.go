package dash

import (
	"context"
	"strings"
	"syscall"

	"k8s.io/klog/v2"
)

// Runner combines all run functions into one interface.
type Runner interface {
	RunE(ctx context.Context,
		splitResult *SplitResult,
		settings ...SettingsFunc) (err error)
	RunWithOutputE(ctx context.Context,
		splitResult *SplitResult,
		settings ...SettingsFunc) (out string, err error)

	/*
		Run(ctx context.Context,
			splitResult SplitResult,
			settings ...SettingsFunc)
		RunWithOutput(ctx context.Context,
			splitResult SplitResult,
			settings ...SettingsFunc) string
	*/

	/*
		RunWithOutputCombined(ctx context.Context,
			splitResult SplitResult,
			settings ...SettingsFunc) string
		RunWithOutputCombinedE(ctx context.Context,
			splitResult SplitResult,
			settings ...SettingsFunc) (out string, err error)
	*/
}

// NewRunner creates the default runner.
// The default log-level is 5 (the highest) to hide debug messages.
func NewRunner(opts ...RunnerOpt) Runner {
	r := &runner{
		logLevel: 5,
	}

	// Apply options.
	for _, opt := range opts {
		opt(r)
	}

	return r
}

// RunnerOpt functions can be passed to NewRunner to configure it for all
// the runners' calls.
type RunnerOpt func(*runner)

type runner struct {
	logLevel klog.Level
}

// LogLevel sets the log-level for log-messages like "Executing ..." .
func LogLevel(level klog.Level) RunnerOpt {
	return func(r *runner) {
		r.logLevel = level
	}
}

/*
func (r *runner) Run(ctx context.Context, splitResult SplitResult,
	settings ...SettingsFunc) {
	if err := r.RunE(ctx, splitResult, settings...); err != nil {
		os.Exit(1)
	}
}
*/

func (r *runner) RunE(ctx context.Context, splitResult *SplitResult,
	settings ...SettingsFunc) (err error) {
	// Check the arguments were parsed without error.
	if splitResult.Err != nil {
		return splitResult.Err
	}

	klog.V(r.logLevel).InfoS("Executing", "command", strings.Join(splitResult.command(), " "))

	setting := extractSettingsFromSlice(settings)

	cmd, cancel, err := createCommand(ctx, splitResult, &setting)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	go func() {
		<-ctx.Done()
		pgid, err := syscall.Getpgid(cmd.Process.Pid)
		if err == nil {
			err = syscall.Kill(-pgid, 15) // note the minus sign
			klog.ErrorS(err, "failed to kill process group",
				"pgid", pgid,
			)
		} else {
			klog.ErrorS(err, "failed to find process group gid")
		}

		err = cmd.Process.Kill()
		klog.ErrorS(err, "failed to kill process the usual way")

		klog.V(5).InfoS("Killed process group",
			"pid", cmd.Process.Pid,
		)
	}()

	defer func() {
		cancel()
	}()

	if err != nil {
		return err
	}

	_, err = runWithSettings(cmd, &setting)
	if err != nil {
		klog.ErrorS(err, "Command failed",
			"name", splitResult.Name,
			"args", splitResult.Args)

		return err
	}

	return nil
}

/*
func (r *runner) RunWithOutput(ctx context.Context, splitResult SplitResult,
	settings ...SettingsFunc) string {
	out, err := r.RunWithOutputE(ctx, splitResult, settings...)
	if err != nil {
		os.Exit(1)
	}
	return out
}
*/

func (r *runner) RunWithOutputE(ctx context.Context, splitResult *SplitResult,
	settings ...SettingsFunc) (out string, err error) {
	// Check the arguments were parsed without error.
	if splitResult.Err != nil {
		return "", splitResult.Err
	}

	klog.V(r.logLevel).InfoS("Executing", "command", strings.Join(splitResult.command(), " "))

	setting := extractSettingsFromSlice(append(settings, withOutput))

	cmd, cancel, err := createCommand(ctx, splitResult, &setting)
	defer func() {
		cancel()
	}()

	if err != nil {
		return "", err
	}

	out, err = runWithSettings(cmd, &setting)
	if err != nil {
		klog.ErrorS(err, "Command failed",
			"name", splitResult.Name,
			"args", splitResult.Args)

		return out, err
	}

	return
}

/*
func (r *runner) RunWithOutputCombined(ctx context.Context, splitResult SplitResult, settings ...SettingsFunc) string {
	panic("not implemented") // TODO: Implement
}

func (r *runner) RunWithOutputCombinedE(ctx context.Context, splitResult SplitResult, settings ...SettingsFunc) (out string, err error) {
	panic("not implemented") // TODO: Implement
}
*/
