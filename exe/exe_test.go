package exe

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"aduu.dev/utils/helper"
	"aduu.dev/utils/helper/testhelper"
)

func TestExecuteWithOutputE(t *testing.T) {
	type args struct {
		s   string
		obj interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"expansion test", args{`echo $HOME`, ""}, os.ExpandEnv("$HOME"), false},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExecuteWithOutputE(tt.args.s, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecuteWithOutputE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExecuteWithOutputE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Execute_UsesWithDir(t *testing.T) {
	tempDir := testhelper.MakeTempDir(t, "execute_UseWithDir")
	defer testhelper.DeleteTempDir(t, tempDir)

	// Creating a so i can try to touch a/b and so won't create a if it did not switch directories correctly
	a := filepath.Join(tempDir, "a")
	b := filepath.Join(tempDir, "a", "b")

	if err := os.Mkdir(a, 0777); err != nil {
		t.Fatalf("failed to create test dir a: %v", err)
	}

	Execute("touch a/b", "", WithDir(tempDir))

	if !helper.DoesPathExist(b) {
		t.Errorf("failed to change directory into %s", tempDir)
		return
	}
	fmt.Println("b does exist:", b)
}
