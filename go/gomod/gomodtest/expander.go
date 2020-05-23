// Package gomodtest includes helpers for accessing files and packages inside the current Go Module.
package gomodtest

import (
	"testing"

	"aduu.dev/utils/expander"
	"aduu.dev/utils/go/gomod"
)

// ExpandPath expands the package path inside tests.
func ExpandPath(t *testing.T, path string) string {
	expandr := makeExpander(t)
	newPath, err := expandr.ExpandPath(path)
	if err != nil {
		t.Fatal(err)
	}
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

func makeExpander(t *testing.T) expander.Expander {
	ws, err := gomod.GetWorkspaceFromWD()
	if err != nil {
		t.Fatal(err)
	}
	return &ws
}
