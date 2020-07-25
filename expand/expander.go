package expand

import (
	"fmt"
)

// Expander is an interface for components which are able to expand path prefixes to absolute paths.
type Expander interface {
	Expand(path string) string
	ExpandFilepath(path string) string
}

var (
	// ErrNoPrefix gets returned if the given path does not have a prefix.
	ErrNoPrefix = fmt.Errorf("path does not have a prefix")
)
