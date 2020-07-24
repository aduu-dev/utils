package projectpath

import "os"

// Projectpath expands the project path variable.
func Projectpath() string {
	return os.ExpandEnv("$PROJECTS")
}