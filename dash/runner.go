package dash

import (
	"context"
	"strings"

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

func NewRunner() Runner {
	return &runner{}
}

type runner struct {
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

	klog.InfoS("Executing", "command", strings.Join(splitResult.command(), " "))

	setting := extractSettingsFromSlice(settings)

	cmd, cancel, err := createCommand(ctx, splitResult, &setting)

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

	klog.InfoS("Executing", "command", strings.Join(splitResult.command(), " "))

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
