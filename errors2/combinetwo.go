package errors2

import "fmt"

// CombineErrors combines errors if both are not nil, otherwise returns the non-nil one if one exists.
func CombineErrors(primary error, secondary error) (err error) {
	switch {
	case primary == nil && secondary == nil:
		err = nil
	case primary == nil && secondary != nil:
		err = secondary
	case primary != nil && secondary == nil:
		err = primary
	case primary != nil && secondary != nil:
		err = fmt.Errorf("secondary failure %#v arose after primary failure: %w", secondary, primary)
	}
	return
}
