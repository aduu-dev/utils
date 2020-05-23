package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"aduu.dev/utils/exe"
	"aduu.dev/utils/helper/testhelper"
)

func deletePath(t *testing.T, path string) {
	err := os.RemoveAll(path)
	assert.NoError(t, err, "failed to delete temporary directory")
}

func TestMoveFile(t *testing.T) {
	type args struct {
		sourcePath string
		createPath string
		destPath   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"same-directory rename", args{"myfile", "myfile", "myfile2"}, false},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			tmp, err := ioutil.TempDir(os.TempDir(), "myprefix")
			assert.NoError(t, err, "failed to create temp directory")
			defer deletePath(t, tmp)

			assert.NoError(t, os.Chdir(tmp), "failed to change to testing directory")

			exe.Execute("touch {{.}}", tt.args.createPath)
			assert.NoError(t, ioutil.WriteFile(tt.args.createPath, []byte("hello"), 0755),
				"failed to write test data to creation path")

			// Doing files not directories.
			//assert.NoError(t, os.MkdirAll(creationPath, 0755), "failed to create directory to move")

			_, err = ioutil.ReadFile(tt.args.createPath)
			assert.NoError(t, err, "createPath should exist before MoveFile()")

			if err := MoveFile(tt.args.sourcePath, tt.args.destPath); (err != nil) != tt.wantErr {
				t.Errorf("MoveFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !DoesPathExist(tt.args.destPath) {
				t.Errorf("MoveFile() file at destPath %s not found", tt.args.destPath)
			}

			if !DoesPathExist(tt.args.createPath) {
				t.Errorf("MoveFile() copy did remove create path %s", tt.args.destPath)
			}

			want, err := ioutil.ReadFile(tt.args.createPath)
			assert.NoError(t, err, "error while reading form create path")

			got, err := ioutil.ReadFile(tt.args.destPath)
			assert.NoError(t, err, "failed to read from destination file path")

			assert.Equal(t, want, got, "source and destination files are not equal")
		})
	}
}

func Test_clearDirectory(t *testing.T) {
	filesToSkipClearing := func() []string {
		return []string{
			".git",
			"LICENSE",
			"README.md",
		}
	}

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		init    func(repoDir string)
	}{
		{
			name: "do not clear non-absolute",
			args: args{
				path: "./",
			},
			wantErr: true,
		},
		{
			name: "do not clear path which does not exist",
			args: args{
				path: "/repo2",
			},
			wantErr: true,
		},
		/*
			{
				name: "do not clear if .git is not inside",
				args: args{
					path: "/repo",
				},
				wantErr: true,
				init: func(repoDir string) {
					if err := os.RemoveAll(filepath.Join(repoDir, ".git")); err != nil {
						panic(fmt.Sprintf("failed to remove repoDir/.git for test: %v", err))
					}
				},
			},
		*/
		{
			name: "clear repository normally",
			args: args{
				path: "/repo",
			},
			wantErr: false,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			tempDir := testhelper.MakeTempDir(t, "ClearDirectory")
			defer testhelper.DeleteTempDir(t, tempDir)

			repoDir := filepath.Join(tempDir, "repo") // Create repo.
			if err := os.Mkdir(repoDir, 0777); err != nil {
				t.Fatal(err)
			}

			creatingFiles := []string{"a", "a/b/c", ".git", "LICENSE", "README.md", "docs"}
			for _, file := range creatingFiles {
				fmt.Println("creating file:", file)
				if err := os.MkdirAll(filepath.Join(repoDir, file), 0777); err != nil {
					t.Fatalf("failed to create test file inside %s: %v", repoDir, err)
				}
			}

			// If starts with slash add tmpdir in front. Can test relative paths this way.
			path := tt.args.path
			if []rune(path)[0] == '/' {
				path = filepath.Join(tempDir, path)
			}

			if tt.init != nil {
				tt.init(path)
			}

			err := ClearDirectory(path, filesToSkipClearing()...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClearDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				// Can't check then since clearing failed anyway.
				return
			}

			wantFiles := []string{".git", "LICENSE", "README.md"}

			// Can do normally.
			if err := testhelper.MustExist(t, repoDir, wantFiles); err != nil {
				t.Errorf("clearDirecotry() failed to clear for test %v: %v", tt.name, err)
			}
		})
	}
}
