package strings2

import (
	"sort"
)

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
//
// Source: https://yourbasic.org/golang/compare-slices/
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// SortedEqual sorts the slices first and then compares like Equal.
func SortedEqual(a, b []string) bool {
	sort.Strings(a)
	sort.Strings(b)

	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
