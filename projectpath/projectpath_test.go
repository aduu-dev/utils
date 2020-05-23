package projectpath

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExpandProjectPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test not expanding if not present", args{path: "/abc"}, "/abc"},
		{
			name: "test not expanding if not present",
			args: args{path: "/abc"},
			want: "/abc",
		},
		{
			name: "test expanding without further path",
			args: args{
				path: "//",
			},
			want: "$HOME/Documents/projects",
		},
		{
			name: "test expanding with further path",
			args: args{
				path: "//abc",
			},
			want: "$HOME/Documents/projects/abc",
		},
		{
			name: "do not expand if not at front",
			args: args{
				path: "/abc//",
			},
			want: "/abc//",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// Expand $HOME.
			tt.want = os.ExpandEnv(tt.want)

			// Use the correct filepath separator for windows & Unix.
			//
			// Its basically a noop on unix.
			//
			// But only fix want if ExpandProjectPth would actually do work.
			if IsProjectPathPresent(tt.args.path) {
				t.Log("name:", tt.name, "replacing:", tt.want)
				tt.want = strings.ReplaceAll(tt.want, "/", string(filepath.Separator))
				t.Log("name:", tt.name, "replaced:", tt.want)
			}

			if got := ExpandProjectPath(tt.args.path); got != tt.want {
				t.Errorf("ExpandProjectPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
