package dash

import (
	"bytes"
	"os"
	"text/template"

	"k8s.io/klog/v2"
)

// TemplateSplitExpand executes the template with obj,
// splits the commands into its parts and then expands env variables.
//
// The expanding is done after the splitting. This is to avoid env variables
// injecting more arguments.
func TemplateSplitExpand(s string, obj interface{}) *SplitResult {
	var templated bytes.Buffer

	if err := templateString(s, obj, &templated); err != nil {
		return &SplitResult{Err: err}
	}

	// Split into arguments.
	split := splitCommand(templated.String())

	// Expand Environment variables for args.
	split = split.expand()

	return split
}

// SplitTemplateExpand splits the argument, then templates
// the result one by one and then expands that one by one.
func SplitTemplateExpand(s string, obj interface{}) *SplitResult {
	// Split into arguments.
	split := splitCommand(s)

	// Template the split.
	split = split.template(obj)

	// Expand the split.
	split = split.expand()

	return split
}

// expand expands each argument with os.Env().
func (s *SplitResult) expand() *SplitResult {
	if s.Err != nil {
		return s
	}

	new := &SplitResult{
		Name: os.ExpandEnv(s.Name),
		Args: make([]string, len(s.Args)),
	}

	// Expand Environment variables for args.
	for i, arg := range s.Args {
		new.Args[i] = os.ExpandEnv(arg)
	}

	return new
}

// template templates each part of s with obj and returns the result.
func (s *SplitResult) template(obj interface{}) *SplitResult {
	if s.Err != nil {
		return s
	}

	new := &SplitResult{Args: make([]string, 0, len(s.Args))}
	var templated bytes.Buffer

	if err := templateString(s.Name, obj, &templated); err != nil {
		return &SplitResult{Err: err}
	}

	new.Name = templated.String()

	for _, arg := range s.Args {
		if err := templateString(arg, obj, &templated); err != nil {
			return &SplitResult{Err: err}
		}

		new.Args = append(new.Args, templated.String())
	}

	return new
}

// template templates s with obj into buf.
func templateString(s string, obj interface{}, buf *bytes.Buffer) (err error) {
	buf.Reset()

	klog.InfoS("string before templating", "string", s)

	// Execute template.
	t, err := template.New("").Parse(s)
	if err != nil {
		return err
	}

	if err := t.Execute(buf, obj); err != nil {
		return err
	}

	return nil
}
