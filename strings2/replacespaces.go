// Package strings2 contains string helpers.
package strings2

import (
	"strings"
	"unicode"
)

// ReplaceSpaces replaces all spaces with the given rune.
func ReplaceSpaces(str string, replacementRune rune) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			// if the character is a space, drop it
			return replacementRune
		}
		// else keep it in the string
		return r
	}, str)
}
