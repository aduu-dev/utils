package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCMD(t *testing.T) {
	cmd := RootCMD()
	buf := new(bytes.Buffer)
	cmd.SetArgs([]string{})
	cmd.SetOutput(buf)

	if _, err := cmd.ExecuteC(); err != nil {
		t.Fatalf("got error fetching project path: %v", err)
	}

	fmt.Println(buf.Len(), buf.String())

	assert.Equal(t, os.ExpandEnv("$HOME/Documents/projects\n"), buf.String())
}
