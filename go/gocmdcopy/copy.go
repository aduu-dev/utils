// Package gocmdcopy copies with the given expander from the given directory into the given directory.
package gocmdcopy

import (
	"fmt"
	"path/filepath"

	"golang.org/x/xerrors"

	"github.com/otiai10/copy"

	"aduu.dev/utils/expander"
)

var (
	// ErrNotAbsolute means that the path wasn't absolute as it was supposed to be.
	ErrNotAbsolute = fmt.Errorf("path is not absolute")
)

// Copy copies from to to. Errors out if to exists already. Expands from and to with the given expander.
func Copy(exp expander.Expander, from string, to string) (err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("copying failed from %s to %s: %w", from, to, err)
		}
	}()

	newFrom, err := exp.ExpandPath(from)
	if err != nil {
		if xerrors.Is(err, expander.ErrNoPrefix) {
			// Do nothing.
		} else {
			return
		}
	} else {
		from = newFrom
	}

	newTo, err := exp.ExpandPath(to)
	if err != nil {
		if xerrors.Is(err, expander.ErrNoPrefix) {
			// Do nothing.
		} else {
			return
		}
	} else {
		to = newTo
	}

	if !filepath.IsAbs(from) {
		return ErrNotAbsolute
	}
	if !filepath.IsAbs(to) {
		return ErrNotAbsolute
	}

	fmt.Println("before:", from, to)

	from, err = filepath.EvalSymlinks(from)
	if err != nil {
		return
	}
	to, err = filepath.EvalSymlinks(to)
	if err != nil {
		return
	}

	fmt.Println("after:", from, to)

	if err = copy.Copy(from, to); err != nil {
		return err
	}
	return nil
}
