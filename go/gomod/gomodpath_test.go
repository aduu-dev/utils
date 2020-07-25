package gomod

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestWorkspace_ExpandGomodPath(t *testing.T) {
	type args struct {
		path string
	}

	tests := []struct {
		name                 string
		workspace            Workspace
		args                 args
		want                 string
		wantErr              bool
		doNotExpandGomodPath bool
	}{
		// OK.
		{
			name: "test simple expanding",
			workspace: Workspace{
				GomodPath: "/hello/world/go.mod",
				Module:    "github.com/hello",
			},
			args: args{
				path: "//",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "test expanding when not only //",
			workspace: Workspace{
				GomodPath: "/hello/world/go.mod",
				Module:    "github.com/hello",
			},
			args: args{
				path: "//abc",
			},
			want:    "/abc",
			wantErr: false,
		},
		{
			name: "do not care if Module stringis zero length",
			workspace: Workspace{
				GomodPath: "/abc/go.mod",
				Module:    "",
			},
			args: args{
				path: "//abc",
			},
			want:    "/abc",
			wantErr: false,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			ws := tt.workspace

			// If I choose a fixed ws.GomodPath here, then I can only make it fixed for one platform.
			// So set it relative and filepath.Abs it here.

			if !tt.doNotExpandGomodPath {
				var err error
				ws.GomodPath, err = filepath.Abs(strings.TrimPrefix(ws.GomodPath, "/"))
				if err != nil {
					t.Fatal(err)
				}

				// Expand want if there is no error.
				if !tt.wantErr {
					tt.want = filepath.Join(filepath.Dir(ws.GomodPath), strings.TrimPrefix(tt.want, "/"))
				}
			}

			// Adapt want to include GomodPath if no error occured.

			got := ws.ExpandFilepath(tt.args.path)
			if got != tt.want {
				t.Errorf("%#v.Expand() =\n got: %v,\nwant: %v", ws, got, tt.want)
			}
		})
	}
}

func TestIsGomodPathPresent(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple //",
			args: args{
				path: "//",
			},
			want: false,
		},
		{
			name: "only one /",
			args: args{
				path: "/",
			},
			want: true,
		},
		{
			name: "three slashes",
			args: args{
				path: "///",
			},
			want: true,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			if got := validatePathForExpansion(tt.args.path); (got != nil) != tt.want {
				t.Errorf("IsGomodPathPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkspace_ExpandPackage(t *testing.T) {
	type fields struct {
		GomodPath string
		Module    string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantS   string
		wantErr bool
	}{
		{
			name: "expand correctly",
			fields: fields{
				GomodPath: "/abc/go.mod",
				Module:    "aduu.dev/hello/world",
			},
			args: args{
				path: "//pkg/internal",
			},
			wantS:   "aduu.dev/hello/world/pkg/internal",
			wantErr: false,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			ws := &Workspace{
				GomodPath: tt.fields.GomodPath,
				Module:    tt.fields.Module,
			}

			// The pre-set GomodPath is not absolute on windows.
			// Let's just expand it here, it is not used anyways.
			var err error
			ws.GomodPath, err = filepath.Abs(strings.TrimPrefix(ws.GomodPath, "/"))
			if err != nil {
				t.Fatal(err)
			}

			gotS, err := ws.ExpandPackage(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workspace.ExpandPackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("Workspace.ExpandPackage() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
