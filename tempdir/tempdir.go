package tempdir

import (
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/xerrors"
)

func MakeTempDir(name string) string {
	tempDir, err := MakeTempDirE(name)
	if err != nil {
		panic(err)
	}
	return tempDir
}

// MakeTempDirE creates a temporary directory with the given name as a prefix.
func MakeTempDirE(name string) (s string, err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("failed to create temporary directory with name %s: %w", name, err)
		}
	}()

	tempDir, err := ioutil.TempDir(os.TempDir(), strings.ReplaceAll(name, " ", "-"))
	if err != nil {
		return
	}
	return tempDir, nil
}

func DeleteTempDir(path string) {
	if err := DeleteTempDirE(path); err != nil {
		panic(err)
	}
}

func DeleteTempDirE(path string) (err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("failed to remove temp dir at %s: %w", err)
		}
	}()

	return os.RemoveAll(path)
}

// DeleteTempDirExistingError deletes the given path and prints corresponding messages if it fails whether an existing error exists or not.
func DeleteTempDirExistingError(path string, existingError *error) (err error) {
	defer func() {
		if err != nil {
			if *existingError != nil {
				err = xerrors.Errorf("cleanup of directory (%s) failed with (%v) while having an error already: %w", path, err, *existingError)
			} else {
				err = xerrors.Errorf("failed to remove temp dir at %s: %w", err)
			}
		}
	}()

	return os.RemoveAll(path)
}

// DeleteTempDirExistingError deletes the given path and prints corresponding messages if it fails whether an existing error exists or not.
func DeleteTempDirExistingErrorNoReturn(path string, existingError *error) {
	cleanupError := os.RemoveAll(path)

	if existingError != nil {
		if *existingError != nil {
			*existingError = xerrors.Errorf("cleanup of directory (%s) failed with (%v) while having an error already: %w", path, cleanupError, *existingError)
		} else {
			*existingError = xerrors.Errorf("failed to remove temp dir at %s: %w", cleanupError)
		}
	} else {
		existingError = &cleanupError
	}
}
