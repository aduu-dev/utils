package gomodtest

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"aduu.dev/utils/helper"
)

// WriteFileJSON writes the given data as JSON to the given path,
// which is being expanded based on the current Go Module.
func WriteFileJSON(t *testing.T, path string, data interface{}) {
	jsn := helper.PrettyJSON(data)

	WriteFile(t, path, jsn)

}

// WriteFile writes the given data to the given path,
// which is being expanded based on the current Go Module.
func WriteFile(t *testing.T, path string, data []byte) {
	expandedPath := ExpandFilepath(t, path)
	if err := ioutil.WriteFile(expandedPath, data, 0755); err != nil {
		t.Fatalf("failed to data to %#v: %v", path, err)
	}
}

// ReadFileJSON reads from the given path data into the given interface.
// Fails fatally if reading the file or unmarshalling into the given data structure fails.
func ReadFileJSON(t *testing.T, path string, data interface{}) {
	readData := ReadFile(t, path)

	if err := json.Unmarshal(readData, &data); err != nil {
		t.Fatalf("failed to unmarshal %#v stored on disk: %v", path, err)
	}
}

// ReadFile reads from the given path data and returns it.
func ReadFile(t *testing.T, path string) []byte {
	expandedPath := ExpandFilepath(t, path)

	readData, err := ioutil.ReadFile(expandedPath)
	if err != nil {
		t.Fatalf("failed to read in testdata from %#v: %v", expandedPath, err)
	}
	return readData
}
