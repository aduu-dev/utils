package helper

import (
	"fmt"
	"os"
	"path/filepath"
)

// DoesPathExist checks whether the given path exists inside the file system.
func DoesPathExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// DoesPathExistErr returns whether the given file or directory exists
func DoesPathExistErr(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// RequireAbsolutePathExisting makes sure the given path is absolute and does actually exist inside the filesystem.
func RequireAbsolutePathExisting(path string) error {
	if !filepath.IsAbs(path) {
		return fmt.Errorf("path is not absolute: %s", path)
	}
	if !DoesPathExist(path) {
		return fmt.Errorf("path does not exist: %s", path)
	}
	return nil
}
