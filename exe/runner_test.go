package exe

/*
type mockRunner struct {
	mock.Mock
}

func (r *mockRunner) Run(cmd *exec.Cmd, name string, settings ...SettingsFunc) {
	panic("implement me")
}

func (r *mockRunner) RunE(cmd *exec.Cmd, name string, settings ...SettingsFunc) (err error) {
	panic("implement me")
}

func (r *mockRunner) RunWithOutput(cmd *exec.Cmd, name string, settings ...SettingsFunc) string {
	panic("implement me")
}

func (r *mockRunner) RunWithOutputE(cmd *exec.Cmd, name string, settings ...SettingsFunc) (out string, err error) {
	panic("implement me")
}

func (r *mockRunner) SplitCommand(s string) []string {
	panic("implement me")
}

func (r *mockRunner) Execute(s string, obj interface{}, settings ...SettingsFunc) *exec.Cmd {
	_ = r.Called(s, obj)
	return nil
}

func (r *mockRunner) ExecuteE(s string, obj interface{}, settings ...SettingsFunc) (*exec.Cmd, error) {
	_ = r.Called(s, obj, settings)
	return nil, nil
}

func (r *mockRunner) ExecuteWithOutput(s string, obj interface{}, settings ...SettingsFunc) string {
	_ = r.Called(s, obj, settings)
}

func (r *mockRunner) ExecuteWithOutputE(s string, obj interface{}, settings ...SettingsFunc) (out string, err error) {
	_ = r.Called(s, obj, settings)
	return "", nil
}

func defaultMockRunner() Runner {
	return &mockRunner{}
}

func TestRunnerMock(t *testing.T) {

}


*/
