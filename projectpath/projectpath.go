// Package projectpath contains functions to expand // to the current project path.
package projectpath

import (
	"os"
	"path/filepath"
	"strings"
)

const projectPath = "$HOME/Documents/projects"
const projectPrefix = "//"

// ExpandProjectPath expands paths with leading //-prefix with the default project path.
func ExpandProjectPath(path string) string {
	if !IsProjectPathPresent(path) {
		return path
	}

	return filepath.Join(ProjectPath(), strings.TrimPrefix(path, projectPrefix))
}

// IsProjectPathPresent says whether the given path is expandable.
func IsProjectPathPresent(path string) bool {
	return strings.HasPrefix(path, projectPrefix)
}

// ProjectPath returns the current project path.
func ProjectPath() string {
	return os.ExpandEnv(projectPath)
}
