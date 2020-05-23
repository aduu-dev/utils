package helper

import (
	"golang.org/x/xerrors"
)

type doNotPrintHelp struct {
	err error
}

func (doNot doNotPrintHelp) Error() string {
	return doNot.err.Error()
}

func (doNot doNotPrintHelp) Unwrap() error {
	return doNot.err
}

func (doNot doNotPrintHelp) As(target interface{}) bool {
	switch target.(type) {
	case doNotPrintHelp, *doNotPrintHelp:
		return true
	default:
		return false
	}
}

// DoNotPrintHelp wraps the error.
func DoNotPrintHelp(err error) error {
	return doNotPrintHelp{
		err: err,
	}
}

// ContainsDoNotPrintHelp says whether the given error says to not print help.
func ContainsDoNotPrintHelp(err error) bool {
	var target doNotPrintHelp
	return xerrors.As(err, &target)
}
