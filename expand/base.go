package expand

import (
	"path"
	"path/filepath"
	"strings"
)

// Base creates a new BaseExpander.
func Base(path string) *BaseExpander {
	return &BaseExpander{
		base: path,
	}
}

const BasePrefix = "//"

// Expand expands paths with leading //-prefix with
// the default project path.
func Expand(base, fp string) string {
	if !HasPrefix(fp) {
		return fp
	}

	return path.Join(base, strings.TrimPrefix(
		strings.TrimSpace(fp), BasePrefix),
	)
}

// ExpandFilepath expands paths with leading //-prefix with
// the default project path.
func ExpandFilepath(base, fp string) string {
	if !HasPrefix(fp) {
		return fp
	}

	return filepath.Join(base, strings.TrimPrefix(
		strings.TrimSpace(fp), BasePrefix),
	)
}

// IsProjectPathPresent says whether the given path is expandable.
func HasPrefix(path string) bool {
	return strings.HasPrefix(path, BasePrefix)
}

var _ Expander = Base(".")

// BaseExpander expands based on the given base.
type BaseExpander struct {
	base string
}

// Base returns the base path BaseExpander adds.
func (b *BaseExpander) Base() string {
	return b.base
}

// Expand replaces BasePrefix at the front with the stored base path.
func (b *BaseExpander) Expand(path string) string {
	return Expand(b.base, path)
}

// ExpandFilepath replaces BasePrefix at the front with the
// stored base filepath.
func (b *BaseExpander) ExpandFilepath(path string) string {
	return ExpandFilepath(b.base, path)
}
