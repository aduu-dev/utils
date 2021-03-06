package logtesting

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"aduu.dev/utils/go/gomod"
	"aduu.dev/utils/tempdir"
	"github.com/stretchr/testify/assert"
	"k8s.io/klog/v2"
)

func TestLogLevel5(t *testing.T) {
	dir := tempdir.MakeTempDir(gomod.DirectCallerFunctionNameWithPackageName())
	tempFile := filepath.Join(dir, "logs.log")

	LogLevel5()

	// The the log-file to read from.
	fs := flag.NewFlagSet("log_testing", flag.ExitOnError)
	klog.InitFlags(fs)

	fs.Set("logtostderr", "false")
	fs.Set("log_file", tempFile)

	// Stop using the logfile later.
	defer func() {
		fs.Set("log_file", "")

		fs.Parse([]string{})
	}()

	if err := fs.Parse([]string{}); err != nil {
		panic("failed to parse flags")
	}

	klog.V(5).InfoS("hello world")
	klog.InfoS("hey du")

	// Make sure the log message is written.
	klog.Flush()

	got, err := ioutil.ReadFile(tempFile)
	if err != nil {
		t.Fatal(err)
	}

	gotString := string(got)
	wantString := "hello world"

	if !assert.True(t,
		strings.Contains(gotString, wantString),
		"log should contain %#v, but is %#v", wantString, gotString,
	) {
		t.Fail()
	}
}
