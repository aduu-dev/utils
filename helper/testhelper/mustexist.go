package testhelper

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MustExist(t *testing.T, inPath string, want []string) error {
	files, err := ioutil.ReadDir(inPath)
	if err != nil {
		return fmt.Errorf("failed to read files in afterCheck after clearing directory: %v", err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	assert.Equal(t, fileNames, want, "should have preserved certain files")

	return nil
}
