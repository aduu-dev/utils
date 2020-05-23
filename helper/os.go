package helper

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// MoveFile copies the given file to the destination. It does not delete the source.
func MoveFile(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("writing to output file failed: %s", err)
	}
	/*
		// The copy was successful, so now delete the original file
		err = os.Remove(sourcePath)
		if err != nil {
			return fmt.Errorf("Failed removing original file: %s", err)
		}
	*/

	return nil
}

// ClearDirectory clears the directory of all files and folders except the listed ones.
// Can only keep files directly inside the folder.
func ClearDirectory(path string, skipFiles ...string) error {
	if !filepath.IsAbs(path) {
		return fmt.Errorf("won't clean non-absolute path %s", path)
	}
	if !DoesPathExist(path) {
		return fmt.Errorf("path to clear does not exist: %s", path)
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read directory")
	}

	clearSkipMap := make(map[string]bool)
	for _, skipFile := range skipFiles {
		clearSkipMap[skipFile] = true
	}

	for _, file := range files {
		// Skip .git-folder.
		if clearSkipMap[file.Name()] {
			continue
		}

		absFile := filepath.Join(path, file.Name())
		if err := os.RemoveAll(absFile); err != nil {
			return fmt.Errorf("failed to remove file %s: %v", absFile, err)
		}
	}

	return nil
}
