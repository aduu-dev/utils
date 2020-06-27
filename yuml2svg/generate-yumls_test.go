package yuml2svg

import (
	"path/filepath"
	"strings"
	"testing"

	"aduu.dev/utils/shell"
	"github.com/stretchr/testify/assert"
)

func Test_svgPath(t *testing.T) {
	type args struct {
		yumlPath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "works",
			args: args{yumlPath: "doc/test.yuml"},
			want: "doc/test.svg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := svgPath(tt.args.yumlPath); got != tt.want {
				t.Errorf("svgPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

var _ shell.Shell = &testShell{}

type testShell struct {
	executedCommands []string
}

func (sh *testShell) RunScript(script string) {
	sh.executedCommands = append(sh.executedCommands, script)
}

func (sh *testShell) Err() error {
	return nil
}

func Test_GenerateYumls(t *testing.T) {
	root := "./testdata"

	sh := &testShell{}
	assert.NoError(t, GenerateYumls(Settings{Dark: true, shell: sh}, root))

	got := sh.executedCommands
	want := []string{
		"cat testdata/.hidden-folder/d.yuml | yuml2svg --dark > testdata/.hidden-folder/d.svg",
		"cat testdata/a.yuml | yuml2svg --dark > testdata/a.svg",
		"cat testdata/b.yuml | yuml2svg --dark > testdata/b.svg",
		"cat testdata/folder/c.yuml | yuml2svg --dark > testdata/folder/c.svg"}

	// Replace / with platform-specific separator.
	for i, cmd := range want {
		want[i] = strings.ReplaceAll(cmd, "/", string(filepath.Separator))
	}

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func TestGenerateYuml(t *testing.T) {
	type args struct {
		settings  Settings
		yumlPath  string
		toSvgPath string
	}

	tests := []struct {
		name         string
		args         args
		wantCommands []string
		wantErr      bool
	}{
		{
			name: "dark mode",
			args: args{
				settings:  Settings{Dark: true},
				yumlPath:  "a.yuml",
				toSvgPath: "b.svg",
			},
			wantCommands: []string{
				"cat a.yuml | yuml2svg --dark > b.svg",
			},
			wantErr: false,
		},
		{
			name: "light mode",
			args: args{
				settings:  Settings{Dark: false},
				yumlPath:  "a.yuml",
				toSvgPath: "b.svg",
			},
			wantCommands: []string{
				"cat a.yuml | yuml2svg > b.svg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := &testShell{}
			tt.args.settings.shell = sh

			if err := GenerateYuml(tt.args.settings, tt.args.yumlPath, tt.args.toSvgPath); (err != nil) != tt.wantErr {
				t.Errorf("GenerateYuml() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := sh.executedCommands
			want := tt.wantCommands

			if !assert.Equal(t, want, got) {
				t.Fail()
			}
		})
	}
}
