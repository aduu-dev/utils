package gocmdcopy

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"

	"aduu.dev/utils/go/gomod"
	"aduu.dev/utils/helper/testhelper"
)

func TestCopy(t *testing.T) {
	dir := testhelper.MakeTempDir(t, strings.Join([]string{"test", t.Name()}, "-"))
	defer testhelper.DeleteTempDir(t, dir)

	testWS, err := gomod.GetWorkspaceFromWD()
	if err != nil {
		t.Fatal(err)
	}

	wsPath, err := testWS.ExpandPath("//go/gocmdcopy/testdata")
	if err != nil {
		t.Fatal(err)
	}

	ws, err := gomod.GetWorkspace(wsPath)
	if err != nil {
		t.Fatal(err)
	}

	if err := Copy(&ws, "//", dir); err != nil {
		t.Fatal(err)
	}

	wantAfter := []string{"a/a.go", "b/b.go", "a.go", "go.mod"}
	sort.Strings(wantAfter)

	// Adapt wanted file paths to the platform separator..
	for i, p := range wantAfter {
		wantAfter[i] = strings.ReplaceAll(p, "/", string(filepath.Separator))
	}

	var foundFiles []string
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return xerrors.Errorf("walk passed error along for path %s: %w", path, err)
		}
		if info.IsDir() {
			return nil
		}

		fmt.Println("Filepath separator:", string(filepath.Separator))
		foundFiles = append(foundFiles, strings.TrimPrefix(strings.TrimPrefix(path, dir), string(filepath.Separator)))
		return nil

	}); err != nil {
		t.Fatal(xerrors.Errorf("failed to retrieve copied over files from %s: %w", dir, err))
	}

	sort.Strings(foundFiles)

	assert.Equal(t, wantAfter, foundFiles)
}
