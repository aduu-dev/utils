// Package errors2 implements methods complementing golang.org/x/xerrors.
package errors2

import (
	"reflect"

	"golang.org/x/xerrors"
)

type any interface {
}

// IsType reports whether any error in err's chain matches target's type.
//
// An error is considered to match a target if it's type is equal to that of the target.
func IsType(err error, target any) bool {
	if target == nil {
		return reflect.TypeOf(err) == reflect.TypeOf(target)
	}

	for {
		if reflect.TypeOf(err) == reflect.TypeOf(target) {
			return true
		}

		if err = xerrors.Unwrap(err); err == nil {
			return false
		}
	}
}
