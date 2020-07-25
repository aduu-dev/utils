package gomod

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/xerrors"

	"aduu.dev/utils/expand"
	"aduu.dev/utils/helper"
)

// Workspace holds the path to the go.mod file and the module namespace.
type Workspace struct {
	GomodPath string
	Module    string
}

var _ expand.Expander = &Workspace{}

// GetWorkspace returns the workspace fromPath is part of.
func GetWorkspace(fromPath string) (ws Workspace, err error) {
	if !helper.DoesPathExist(fromPath) {
		return Workspace{}, xerrors.Errorf("path %s does not exist", fromPath)
	}

	fromPath, err = filepath.Abs(fromPath)
	if err != nil {
		return
	}

	gomodPath, module, err := determineGomodPathAndModule(fromPath)
	if err != nil {
		return Workspace{}, err
	}

	return Workspace{
		GomodPath: gomodPath,
		Module:    module,
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

// determineModule returns the module path from the gomod file text.
// If it cannot find a module path, it returns an empty string.
// It is tolerant of unrelated problems in the go.mod file.
func determineModule(mod []byte) string {
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

// determineGomodPathAndModule determines gomod path and module starting from the given path.
func determineGomodPathAndModule(from string) (gomodPath string, module string, err error) {
	gomodPath, err = findGoModFrom(from)
	if err != nil {
		return "", "", fmt.Errorf("failed find go.mod-file starting from %s: %w", from, err)
	}

	moduleContent, err := ioutil.ReadFile(gomodPath)
	if err != nil {
		return "", "", fmt.Errorf("failed to read go.mod-file: %v", err)
	}

	module = determineModule(moduleContent)

	return
}
