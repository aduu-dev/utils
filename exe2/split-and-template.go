package exe2

import (
	"bytes"
	"os"
	"text/template"
)

// TemplateSplitExpand executes the template with obj,
// splits the commands into its parts and then expands env variables.
//
// The expanding is done after the splitting. This is to avoid env variables
// injecting more arguments.
func TemplateSplitExpand(s string, obj interface{}) SplitResult {
	// Execute template.
	t, err := template.New("").Parse(s)
	if err != nil {
		return SplitResult{Err: err}
	}

	var templated bytes.Buffer

	if err := t.Execute(&templated, obj); err != nil {
		return SplitResult{Err: err}
	}

	// Split into arguments.
	split := splitCommand(templated.String())

	// Expand Environment variables for args.
	for i, arg := range split.Args {
		split.Args[i] = os.ExpandEnv(arg)
	}

	// Expand Environment variables for name.
	split.Name = os.ExpandEnv(split.Name)

	return split
}
