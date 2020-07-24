// Package packages includes methods to help dealing with go packages.
package packages // import "aduu.dev/utils/packages"

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"aduu.dev/utils/dash"
)

// From path returns the package name starting from the given path.
func From(path string) (string, error) {
	runner := dash.NewRunner()

	out, err := runner.RunWithOutputE(
		context.Background(),
		dash.TemplateSplitExpand(
			`go list`,
			"",
		), dash.WithDir(path))
	if err != nil {
		return "", err
	}

	return out, nil

	/*
		ws, err := gomod.GetWorkspaceFromWD()
		if err != nil {
			return "", xerrors.Errorf("failed to find go.mod to determine current package: %w", err)
		}

		currentDir := path

		root := filepath.Dir(ws.GomodPath)

		// Add module path in front, then add current file with root removed.
		return ws.Module + "/" + strings.TrimPrefix(currentDir, root+"/"), nil
	*/
}

func evalSymlinks(path string) (string, error) {
	fromTemp, err := filepath.EvalSymlinks(path)
	if err != nil {
		return "", fmt.Errorf("failed to eval symlinks for path %s: %v", path, err)
	}
	return fromTemp, nil
}

// ListPackages lists all packages inside the given path and all sub-directories.
func ListPackages(r dash.Runner, fromPath string) ([]string, error) {
	out, err := r.RunWithOutputE(
		context.Background(),
		dash.TemplateSplitExpand(
			`go list ./...`,
			"",
		), dash.WithDir(fromPath))
	if err != nil {
		return nil, err
	}

	return strings.Split(out, "\n"), nil

	/*
		// Make fromPath absolute so removal of prefix is simpler later.
		fromPath, err := filepath.Abs(fromPath)
		if err != nil {
			return nil, err
		}

		ws, err := gomod.GetWorkspace(fromPath)
		if err != nil {
			return nil, fmt.Errorf("failed to list packages: %v", err)
		}

		// Evaluate symlinks since Walk possibly can't handle it.
		fromPath, err = filepath.EvalSymlinks(fromPath)
		if err != nil {
			return nil, err
		}

		root := filepath.Dir(ws.GomodPath)

		// Evaluate symlinks since Walk possibly can't handle it.
		root, err = evalSymlinks(root)
		if err != nil {
			return nil, err
		}

		var packages []string
		if err := filepath.Walk(fromPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return xerrors.Errorf("got error from filepath.Walk: %w", err)
			}

			fmt.Println("path:", path)

			if !info.IsDir() {
				return nil
			}

			filename := filepath.Base(path)
			if filename == "testdata" {
				return filepath.SkipDir
			}

			// Trim root and a path separator at the end.
			trimmedPath := strings.TrimPrefix(strings.TrimPrefix(path, root), string(filepath.Separator))

			var pkg string
			if trimmedPath == "" {
				pkg = ws.Module
			} else {
				pkg = ws.Module + "/" + trimmedPath
			}
			packages = append(packages, pkg)
			return nil
		}); err != nil {
			return nil, fmt.Errorf("failed to list files: %v", err)
		}

		return packages, nil
	*/
}
