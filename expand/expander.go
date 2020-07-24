package expand

import (
	"fmt"
)

// Expander is an interface for components which are able to expand path prefixes to absolute paths.
type Expander interface {
	ExpandPath(path string) (string, error)
	IsExpandable(path string) error
}

var (
	// ErrNoPrefix gets returned if the given path does not have a prefix.
	ErrNoPrefix = fmt.Errorf("path does not have a prefix")
)
