package helper

import (
	"os"
	"path/filepath"
	"testing"

	"aduu.dev/utils/helper/testhelper"
)

func TestRequireAbsolutePathExisting(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name              string
		args              args
		doNotMakeAbsolute bool
		wantErr           bool
	}{
		{
			name: "does not exist",
			args: args{
				path: "non-existent",
			},
			wantErr: true,
		},
		{
			name: "is not absolute",
			args: args{
				path: "non-existent",
			},
			doNotMakeAbsolute: true,
			wantErr:           true,
		},
		{
			name: "does exist",
			args: args{
				path: "abc",
			},
			wantErr: false,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			dir := testhelper.MakeTempDir(t, tt.name)
			defer testhelper.DeleteTempDir(t, dir)

			if err := os.Mkdir(filepath.Join(dir, "abc"), 0777); err != nil {
				t.Fatalf("failed to create abc helper directory: %v", err)
			}

			if !tt.doNotMakeAbsolute {
				tt.args.path = filepath.Join(dir, tt.args.path)
			}

			if err := RequireAbsolutePathExisting(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("RequireAbsolutePathExisting() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoesPathExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check existence of file",
			args: args{
				path: "file",
			},
			want: true,
		},
		{
			name: "check existence of directory",
			args: args{
				path: "dir",
			},
			want: true,
		},
		{
			name: "non-existent",
			args: args{
				path: "abc",
			},
			want: false,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			dir := testhelper.MakeTempDir(t, tt.name)
			defer testhelper.DeleteTempDir(t, dir)

			if err := os.Mkdir(filepath.Join(dir, "dir"), 0777); err != nil {
				t.Fatalf("failed to create dir - helper directory: %v", err)
			}
			if file, err := os.Create(filepath.Join(dir, "file")); err != nil {
				t.Fatalf("failed to create file - helper file: %v", err)
			} else {
				_ = file.Close()
			}

			tt.args.path = filepath.Join(dir, tt.args.path)

			if got := DoesPathExist(tt.args.path); got != tt.want {
				t.Errorf("DoesPathExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
