package gomod

import (
	"testing"
)

func TestWorkspace_ExpandGomodPath(t *testing.T) {
	type args struct {
		path string
	}

	workspace := Workspace{
		GomodPath: "/hello/world/go.mod",
		Module:    "github.com/hello",
	}

	tests := []struct {
		name      string
		workspace Workspace
		args      args
		want      string
		wantErr   bool
	}{
		{
			name: "error if GomodPath does not have /go.mod at the end",
			workspace: Workspace{
				GomodPath: "/abc",
				Module:    "",
			},
			args: args{
				path: "//abc",
			},
			want:    "",
			wantErr: true,
		},
		{
			name:      "test error if not present",
			workspace: workspace,
			args: args{
				path: "",
			},
			want:    "",
			wantErr: true,
		},
		{
			name:      "test simple expanding",
			workspace: workspace,
			args: args{
				path: "//",
			},
			want:    "/hello/world",
			wantErr: false,
		},
		{
			name:      "test expanding when not only //",
			workspace: workspace,
			args: args{
				path: "//abc",
			},
			want:    "/hello/world/abc",
			wantErr: false,
		},
		{
			name: "test error if GomodPath is == /go.mod",
			workspace: Workspace{
				GomodPath: "/go.mod",
				Module:    "abc",
			},
			args: args{
				path: "//abc",
			},
			want:    "",
			wantErr: true,
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
			want:    "/abc/abc",
			wantErr: false,
		},
		{
			name: "error out if GomodPath is not absolute length",
			workspace: Workspace{
				GomodPath: "abc/go.mod",
				Module:    "",
			},
			args: args{
				path: "//abc",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "error if GomodPath has slash at the end",
			workspace: Workspace{
				GomodPath: "/abc/go.mod/",
				Module:    "",
			},
			args: args{
				path: "//abc",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "error if path starts with more than two slashes at the beginning",
			workspace: Workspace{
				GomodPath: "/abc/go.mod",
				Module:    "",
			},
			args: args{
				path: "///abc",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			ws := tt.workspace
			got, err := ws.ExpandPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workspace.ExpandPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Workspace.ExpandPath() = %v, want %v", got, tt.want)
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
