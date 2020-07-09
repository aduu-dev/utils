package yuml2svg

import (
	"path/filepath"
	"strings"
	"testing"

	"aduu.dev/utils/dash"
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

func Test_GenerateYumls(t *testing.T) {
	root := "./testdata"

	r := dash.NewTestRunner()
	assert.NoError(t, GenerateYumls(Settings{Dark: true, r: r}, root))

	got := r.Commands()
	want := []dash.Command{
		{
			Command: []string{"yuml2svg", "--dark"},
			Setting: dash.ExecuteSetting{
				StdinFile:  "testdata/.hidden-folder/d.yuml",
				StdoutFile: "testdata/.hidden-folder/d.svg",
			},
		},
		{
			Command: []string{"yuml2svg", "--dark"},
			Setting: dash.ExecuteSetting{
				StdinFile:  "testdata/a.yuml",
				StdoutFile: "testdata/a.svg",
			},
		},
		{
			Command: []string{"yuml2svg", "--dark"},
			Setting: dash.ExecuteSetting{
				StdinFile:  "testdata/b.yuml",
				StdoutFile: "testdata/b.svg",
			},
		},
		{
			Command: []string{"yuml2svg", "--dark"},
			Setting: dash.ExecuteSetting{
				StdinFile:  "testdata/folder/c.yuml",
				StdoutFile: "testdata/folder/c.svg",
			},
		},
	}

	// Replace / with platform-specific separator.
	for i, cmd := range want {
		for j, arg := range cmd.Command {
			cmd.Command[j] =
				strings.ReplaceAll(arg, "/", string(filepath.Separator))
		}

		cmd.Setting.StdoutFile =
			strings.ReplaceAll(
				cmd.Setting.StdoutFile, "/", string(filepath.Separator),
			)

		cmd.Setting.StderrFile =
			strings.ReplaceAll(
				cmd.Setting.StderrFile, "/", string(filepath.Separator),
			)

		cmd.Setting.StdinFile =
			strings.ReplaceAll(
				cmd.Setting.StdinFile, "/", string(filepath.Separator),
			)

		want[i] = cmd
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
		name        string
		args        args
		wantCommand dash.Command
		wantErr     bool
	}{
		{
			name: "dark mode",
			args: args{
				settings:  Settings{Dark: true},
				yumlPath:  "a.yuml",
				toSvgPath: "b.svg",
			},
			wantCommand: dash.Command{
				Command: []string{"yuml2svg", "--dark"},
				Setting: dash.ExecuteSetting{
					StdinFile:  "a.yuml",
					StdoutFile: "b.svg",
				},
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
			wantCommand: dash.Command{
				Command: []string{"yuml2svg"},
				Setting: dash.ExecuteSetting{
					StdinFile:  "a.yuml",
					StdoutFile: "b.svg",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := dash.NewTestRunner()
			tt.args.settings.r = r

			if err := GenerateYuml(tt.args.settings, tt.args.yumlPath, tt.args.toSvgPath); (err != nil) != tt.wantErr {
				t.Errorf("GenerateYuml() error = %v, wantErr %v", err, tt.wantErr)
			}

			got := r.Commands()
			want := []dash.Command{tt.wantCommand}

			if !assert.Equal(t, want, got) {
				t.Fail()
			}
		})
	}
}
