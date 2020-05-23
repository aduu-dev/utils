package gomod

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"golang.org/x/xerrors"
)

const gomodTemplate = `module abc.test

go 1.12

require (
	{{.}}
)`

// Make constructs a go.mod file containing the given packages.
func Make(packageLines []string) (string, error) {
	if len(packageLines) == 0 {
		return "", fmt.Errorf("make requires package lines to construct the file")
	}

	tmpl, err := template.New("gomodTemplate").Parse(gomodTemplate)
	if err != nil {
		return "", xerrors.Errorf("failed to make go.mod file: %w", err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, strings.Join(packageLines, "\n	")); err != nil {
		return "", xerrors.Errorf("failed to make go.mod file: %w", err)
	}
	return buf.String(), nil
}
