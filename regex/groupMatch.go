// Package regex contains regexp helper methods for group matches.
package regex

import (
	"regexp"

	"golang.org/x/xerrors"
)

// GroupMatch returns map if it found a match.
func GroupMatch(data string, regexString string) (match bool, m map[string]string, err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("group match failed: %w", err)
		}
	}()

	r, err := regexp.Compile(regexString)
	if err != nil {
		return
	}

	matches := r.FindStringSubmatch(data)
	if matches == nil {
		return
	}

	m = make(map[string]string)
	match = true

	names := r.SubexpNames()
	for i := range matches {
		m[names[i]] = matches[i]
	}
	return
}

// GroupMatchAll returns array of maps if it found a match.
func GroupMatchAll(data string, regexString string) (match bool, maps []map[string]string, err error) {
	defer func() {
		if err != nil {
			err = xerrors.Errorf("group match failed: %w", err)
		}
	}()

	r, err := regexp.Compile(regexString)
	if err != nil {
		return
	}

	matches := r.FindAllStringSubmatch(data, -1)
	if matches == nil {
		return
	}

	match = true

	names := r.SubexpNames()

	for _, matchArray := range matches {
		m := make(map[string]string)
		for i := range matchArray {
			m[names[i]] = matchArray[i]
		}
		maps = append(maps, m)
	}

	return
}
