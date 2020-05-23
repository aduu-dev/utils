package errors2

import (
	"fmt"
)

var (
	ErrPathDoesNotExist = fmt.Errorf("path does not exist")
)
