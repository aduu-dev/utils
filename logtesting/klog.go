// Package logtesting can set the logLevel inside tests to show debug messages printed with klog.
package logtesting

import (
	"flag"

	"k8s.io/klog/v2"
)

// LogLevel5 sets the log level to 5 for testing.
func LogLevel5() {
	fs := flag.NewFlagSet("log", flag.ExitOnError)
	klog.InitFlags(fs)

	if err := fs.Parse([]string{"-v", "5"}); err != nil {
		panic("failed to parse flags")
	}
}
