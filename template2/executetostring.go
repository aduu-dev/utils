package template2

import (
	"bytes"
	"html/template"

	"golang.org/x/xerrors"
)

// ExecuteToString executes the given query into a string buffer.
func ExecuteToStringE(tmpl string, arg interface{}) (out string, err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("failed to execute template %#v with %#v: %w", tmpl, arg, err)
		}
	}()

	var buf bytes.Buffer
	t, err := template.New("temporary-template").Parse(tmpl)
	if err != nil {
		return
	}

	if err = t.Execute(&buf, arg); err != nil {
		return
	}

	out = buf.String()
	return
}

// ExecuteToString executes the given query into a string buffer.
func ExecuteToString(tmpl string, arg interface{}) (out string) {
	out, err := ExecuteToStringE(tmpl, arg)
	if err != nil {
		panic(err)
	}
	return out
}
