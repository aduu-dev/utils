package gomod

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"aduu.dev/utils/helper/testhelper"
)

func TestGetWorkspace_found(t *testing.T) {
	got, err := GetWorkspace(filepath.Join("testdata", "a"))
	assert.NoError(t, err)

	// Bazel turns the paths weird.
	got.GomodPath = ""

	assert.Equal(t, &Workspace{
		GomodPath: "",
		Module:    "aduu.dev",
	}, &got)
}

func TestGetWorkspace_errorIfNotFound(t *testing.T) {
	dir := testhelper.MakeTempDir(t, strings.Join([]string{"test", t.Name()}, "-"))
	defer testhelper.DeleteTempDir(t, dir)

	_, err := GetWorkspace(dir)
	assert.Error(t, err, "should have not found go.mod in a standard environment"+
		"where tmp directory climbing does n have a parent go.mod.")
}

func TestGetWorkspace_errorGivenPathDoesNotExist(t *testing.T) {
	path := filepath.Join(DirectCallerPackagePath(), "testdata", "a", "non-existent-dir")

	_, err := GetWorkspace(path)
	assert.Error(t, err)
}
