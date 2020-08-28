package dmock

import (
	"context"

	"aduu.dev/utils/dash"
	"github.com/stretchr/testify/mock"
)

// New creates a new MockRunner.
func New() *MockRunner {
	return &MockRunner{}
}

type WithDir struct {
	Dir string
}

// MockRunner is used to mock shell calls.
type MockRunner struct {
	mock.Mock
}

func applySettings(settings []dash.SettingsFunc) *dash.ExecuteSetting {
	res := &dash.ExecuteSetting{}

	for _, set := range settings {
		set(res)
	}

	return res
}

func (m *MockRunner) RunE(ctx context.Context,
	splitResult *dash.SplitResult,
	settings ...dash.SettingsFunc,
) (err error) {
	args := m.Called(ctx, splitResult, applySettings(settings))

	return args.Error(0)
}

func (m *MockRunner) RunWithOutputE(ctx context.Context,
	splitResult *dash.SplitResult,
	settings ...dash.SettingsFunc,
) (out string, err error) {
	args := m.Called(ctx, splitResult, applySettings(settings))

	return args.String(0), args.Error(1)
}
