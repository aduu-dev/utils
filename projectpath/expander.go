package projectpath

import (
	"fmt"

	"aduu.dev/utils/expander"
)

// Expander returns an expander based on the current projectpath.
func Expander() expander.Expander {
	return &ppathExpander{}
}

type ppathExpander struct {
}

func (e *ppathExpander) ExpandPath(path string) (string, error) {
	if err := e.IsExpandable(path); err != nil {
		return "", err
	}
	return ExpandProjectPath(path), nil
}

func (e *ppathExpander) ExpandPackage(path string) (string, error) {
	panic("implement me")
}

func (e *ppathExpander) IsExpandable(path string) error {
	if IsProjectPathPresent(path) {
		return nil
	}
	return fmt.Errorf("path does not contain the projectpath-prefix //")
}
