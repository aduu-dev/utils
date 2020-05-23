package gomod

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/xerrors"

	"aduu.dev/utils/helper"
)

// Workspace holds the path to the go.mod file and the module namespace.
type Workspace struct {
	GomodPath string
	Module    string
}

// GetWorkspace returns the workspace fromPath is part of.
func GetWorkspace(fromPath string) (ws Workspace, err error) {
	if !helper.DoesPathExist(fromPath) {
		return Workspace{}, xerrors.Errorf("path %s does not exist", fromPath)
	}

	fromPath, err = filepath.Abs(fromPath)
	if err != nil {
		return
	}

	gomodPath, moduleContent, err := modPathFrom(fromPath)
	if err != nil {
		return Workspace{}, err
	}

	return Workspace{
		GomodPath: gomodPath,
		Module:    moduleContent,
	}, nil
}

// GetWorkspaceFromWD returns the workspace from the working directory.
func GetWorkspaceFromWD() (Workspace, error) {
	wd, err := os.Getwd()
	if err != nil {
		return Workspace{}, xerrors.Errorf("failed to get wd so did not find go.mod path: %v", err)
	}
	return GetWorkspace(wd)
}

func findGoModFrom(from string) (string, error) {
	possibleGomod := filepath.Join(from, "go.mod")
	if helper.DoesPathExist(possibleGomod) {
		return possibleGomod, nil
	}

	nextToTry := filepath.Dir(from)
	if nextToTry == from {
		return "", fmt.Errorf("no go.mod in path")
	}

	return findGoModFrom(nextToTry)
}

func slashSlash() []byte {
	return []byte("//")
}

func moduleString() []byte {
	return []byte("module")
}

// Source: https://github.com/golang/go/blob/master/src/cmd/go/internal/modfile/read.go#L837

// modulePath returns the module path from the gomod file text.
// If it cannot find a module path, it returns an empty string.
// It is tolerant of unrelated problems in the go.mod file.
func modulePath(mod []byte) string {
	for len(mod) > 0 {
		line := mod
		mod = nil
		if i := bytes.IndexByte(line, '\n'); i >= 0 {
			line, mod = line[:i], line[i+1:]
		}
		if i := bytes.Index(line, slashSlash()); i >= 0 {
			line = line[:i]
		}
		line = bytes.TrimSpace(line)
		if !bytes.HasPrefix(line, moduleString()) {
			continue
		}
		line = line[len(moduleString()):]
		n := len(line)
		line = bytes.TrimSpace(line)
		if len(line) == n || len(line) == 0 {
			continue
		}

		if line[0] == '"' || line[0] == '`' {
			p, err := strconv.Unquote(string(line))
			if err != nil {
				return "" // malformed quoted string or multiline module path
			}
			return p
		}

		return string(line)
	}
	return "" // missing module path
}

// modPathFrom determines gomod path and modpath starting from the given path.
func modPathFrom(from string) (string, string, error) {
	gomodPath, err := findGoModFrom(from)
	if err != nil {
		return "", "", fmt.Errorf("failed find go.mod-file starting from %s: %v", from, err)
	}

	moduleContent, err := ioutil.ReadFile(gomodPath)
	if err != nil {
		return "", "", fmt.Errorf("failed to read go.mod-file: %v", err)
	}
	return gomodPath, modulePath(moduleContent), nil
}
