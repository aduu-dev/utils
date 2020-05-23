package helper

import (
	"fmt"
	"os"
)

// Getwd returns the current working directory. Panics if that fails.
func Getwd() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}

// Printwd prints the working directory.
func Printwd() {
	fmt.Println(Getwd())
}
