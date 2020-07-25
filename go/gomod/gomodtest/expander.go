// Package gomodtest includes helpers for accessing files and packages inside the current Go Module.
package gomodtest

import (
	"testing"

	"aduu.dev/utils/go/gomod"
)

// ExpandPath expands the package path inside tests.
func ExpandFilepath(t *testing.T, path string) string {
	expandr := makeExpander(t)
	newPath := expandr.ExpandFilepath(path)

	return newPath
}

// ExpandPackage expands the package name inside tests.
func ExpandPackage(t *testing.T, path string) string {
	expandr := makeExpander(t)
	newPath, err := expandr.ExpandPackage(path)

	if err != nil {
		t.Fatal(err)
	}

	return newPath
}

func makeExpander(t *testing.T) *gomod.Workspace {
	ws, err := gomod.GetWorkspaceFromWD()
	if err != nil {
		t.Fatal(err)
	}

	return &ws
}
