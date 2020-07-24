// Package gocmdcopy copies with the given expander from the given directory into the given directory.
package gocmdcopy

import (
	"fmt"
	"path/filepath"

	"golang.org/x/xerrors"

	"github.com/otiai10/copy"

	"aduu.dev/utils/expand"
)

var (
	// ErrNotAbsolute means that the path wasn't absolute as it was supposed to be.
	ErrNotAbsolute = fmt.Errorf("path is not absolute")
)

// Copy copies from to to. Errors out if to exists already. Expands from and to with the given expander.
func Copy(exp *expand.BaseExpander, from string, to string) (err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("copying failed from %s to %s: %w", from, to, err)
		}
	}()

	from = exp.ExpandFilepath(from)
	to = exp.ExpandFilepath(to)

	if !filepath.IsAbs(from) {
		return ErrNotAbsolute
	}
	if !filepath.IsAbs(to) {
		return ErrNotAbsolute
	}

	from, err = filepath.EvalSymlinks(from)
	if err != nil {
		return
	}
	to, err = filepath.EvalSymlinks(to)
	if err != nil {
		return
	}

	if err = copy.Copy(from, to); err != nil {
		return err
	}
	return nil
}
