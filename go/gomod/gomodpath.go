package gomod

import (
	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/xerrors"

	"aduu.dev/utils/expander"
)

const projectPrefix = "//"

// ExpandPath expands paths with leading //-prefix with the folder the nearest go.mod resides in.
func (ws *Workspace) ExpandPath(path string) (s string, err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("failed to expand with Workspace %#v: %w", ws, err)
		}
	}()

	if err = ws.validateExpansion(path); err != nil {
		return
	}
	gomodDir := ws.GomodPathDirectory()

	// Trim project prefix.
	path = strings.TrimPrefix(path, projectPrefix)

	return filepath.Join(gomodDir, path), nil
}

func (ws Workspace) validateExpansion(path string) (err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("workspace is unable to be used for expansion (%#v): %w", ws, err)
		}
	}()

	if strings.HasSuffix(ws.GomodPath, "/") {
		return fmt.Errorf("GomodPath is not allowed to have suffix /")
	}

	if filepath.Base(ws.GomodPath) != "go.mod" {
		return fmt.Errorf("workspace.GomodPath must have /go.mod-suffix")
	}

	if err = validatePathForExpansion(path); err != nil {
		return
	}

	gomodDir := ws.GomodPathDirectory()
	if gomodDir == "/" {
		return fmt.Errorf("gomod dir too short: %s", gomodDir)
	}
	if !filepath.IsAbs(gomodDir) {
		return fmt.Errorf("gomod dir %#v is not absolute", gomodDir)
	}

	return nil
}

// IsGomodPathPresent says whether the given path is expandable.
func validatePathForExpansion(path string) error {
	if !strings.HasPrefix(path, projectPrefix) {
		return expander.ErrNoPrefix
	}

	runes := []rune(path)
	if len(runes) > 2 {
		// Check its only two slashes in front and not more.
		if runes[2] == '/' {
			return fmt.Errorf("after removing // from path, it still has more than zero / at the beginning: %s", path)
		}
	}

	return nil
}

// GomodPathDirectory returns the directory where the nearest go.mod resides..
func (ws *Workspace) GomodPathDirectory() string {
	return filepath.Dir(ws.GomodPath)
}

// ExpandPackage expands the partial package path and expands it with the path located in the go.mod file.
func (ws *Workspace) ExpandPackage(path string) (s string, err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("failed to expand to full package for (%s): %w", path, err)
		}
	}()

	if err = ws.validateExpansion(path); err != nil {
		return
	}

	return ws.Module + "/" + strings.TrimPrefix(path, projectPrefix), nil
}

// IsExpandable checks whether expanding the given path with the given Workspace would result in an error.
func (ws *Workspace) IsExpandable(path string) error {
	return ws.validateExpansion(path)
}
