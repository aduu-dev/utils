package expand_test

import (
	"path/filepath"
	"testing"

	"aduu.dev/utils/expand"
)

func TestExpand(t *testing.T) {
	type args struct {
		base string
		path string
	}
	tests := []struct {
		name         string
		args         args
		wantPath     string
		wantFilepath string
	}{
		{
			name:     "test not expanding if not present",
			args:     args{base: "", path: "/abc"},
			wantPath: "/abc",
			wantFilepath: "/abc",
		},
		{
			name:         "test not expanding if not present",
			args:         args{base: "", path: "/abc"},
			wantPath:     "/abc",
			wantFilepath: "/abc",
		},
		{
			name: "test expanding without further path",
			args: args{
				base: "hello",
				path: "//",
			},
			wantPath:     "hello",
			wantFilepath: "hello",
		},
		{
			name: "test expanding with further path",
			args: args{
				base: "abc",
				path: "//def",
			},
			wantPath:     "abc/def",
			wantFilepath: "abc" + string(filepath.Separator) + "def",
		},
		{
			name: "do not expand if not at front",
			args: args{
				base: "abc",
				path: "/def//",
			},
			wantPath:     "/def//",
			wantFilepath: "/def//",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			// Use the correct filepath separator for windows & Unix.
			//
			// Its basically a noop on unix.
			//
			// But only fix want if ExpandProjectPth would actually do work.
			/*
				if IsProjectPathPresent(tt.args.path) {
					t.Log("name:", tt.name, "replacing:", tt.want)
					tt.want = strings.ReplaceAll(tt.want, "/", string(filepath.Separator))
					t.Log("name:", tt.name, "replaced:", tt.want)
				}
			*/

			if gotPath := expand.Expand(tt.args.base, tt.args.path); gotPath != tt.wantPath {
				t.Errorf("expand.Expand(%#v, %#v) = %#v, want %#v", tt.args.base, tt.args.path, gotPath, tt.wantPath)
			}

			if gotFilepath := expand.ExpandFilepath(tt.args.base, tt.args.path); gotFilepath != tt.wantFilepath {
				t.Errorf("expand.ExpandFilepath(%#v, %#v) = %#v, want %#v",
					tt.args.base, tt.args.path, gotFilepath, tt.wantFilepath)
			}

			exp := expand.Base(tt.args.base)

			if gotPath := exp.Expand(tt.args.path); gotPath != tt.wantPath {
				t.Errorf("exp.Expand(%#v) = %#v, want %#v",
					tt.args.path, gotPath, tt.wantPath)
			}

			if gotFilepath := exp.ExpandFilepath(tt.args.path); gotFilepath != tt.wantFilepath {
				t.Errorf("exp.ExpandFilepath(%#v) = %#v, want %#v",
					tt.args.path, gotFilepath, tt.wantFilepath)
			}
		})
	}
}
