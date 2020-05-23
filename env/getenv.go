package env

import (
	"fmt"
	"os"
)

// GetenvRequired panics if the environment variable is not set.
func GetenvRequired(key string) string {
	value, err := GetenvE(key)
	if err != nil {
		panic(err)
	}
	return value
}

// GetenvE returns an error if the value is zero length, otherwise returns the environment variable.
func GetenvE(key string) (value string, err error) {
	value = os.Getenv(key)
	if len(value) == 0 {
		err = fmt.Errorf("%#v env key is not set", key)
	}
	return
}
