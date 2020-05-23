// Package fmt2 puts some fmt-Package methods into an interface.
package fmt2

import (
	"fmt"
)

// Printer includes methods which mirror the fmt package.
type Printer interface {
	Println(i ...interface{})
	Printf(format string, i ...interface{})
}

// DefaultPrinter prints with fmt.
func DefaultPrinter() Printer {
	return &defaultPrinter{}
}

type defaultPrinter struct {
}

func (d *defaultPrinter) Println(i ...interface{}) {
	fmt.Println(i...)
}

func (d *defaultPrinter) Printf(format string, i ...interface{}) {
	fmt.Printf(format, i...)
}
