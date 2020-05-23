package difftest

import (
	"testing"

	"aduu.dev/utils/exe"
)

// Diff prints differences between the two given folders.
func Diff(t *testing.T, got string, want string) {
	opts := struct {
		GotDir  string
		WantDir string
	}{
		GotDir:  got,
		WantDir: want,
	}

	if _, err := exe.ExecuteE(`diff -r {{.WantDir}} {{.GotDir}}`, opts); err != nil {
		// Do not delete tempdir in case diff failed.
		t.Fatalf("%s() => err = %v", t.Name(), err)
	}
}
 