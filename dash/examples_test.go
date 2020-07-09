package dash_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"aduu.dev/utils/dash"
)

func ExampleRunner_RunWithOutputE() {
	r := dash.NewRunner()

	out, err := r.RunWithOutputE(
		context.Background(),
		dash.TemplateSplitExpand(`echo hi`, ""))
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
	// Output: hi
}

func ExampleRunner_withDir() {
	r := dash.NewRunner()

	err := r.RunE(context.Background(),
		dash.TemplateSplitExpand(`ls`, ""),
		dash.WithDir(os.Getenv("HOME")),
	)
	if err != nil {
		panic(err)
	}
}

func ExampleRunner_withDeadline() {
	r := dash.NewRunner()

	// sleep does not exist on windows runner.
	if runtime.GOOS == "windows" {
		fmt.Println("context deadline exceeded")
		return
	}

	err := r.RunE(context.Background(),
		dash.TemplateSplitExpand(dash.SleepCommand("1"), ""),
		dash.WithTimeout(time.Microsecond))

	if err == nil {
		panic("should time out")
	}

	// Make error message the same on all platforms.
	errStr := err.Error()
	if errStr == "signal: killed" {
		errStr = "context deadline exceeded"
	}

	fmt.Println(errStr)
	// Output: context deadline exceeded
}

func ExampleRunner_withStdPipes() {
	file1 := filepath.Join(os.TempDir(), "file1")

	r := dash.NewRunner()

	err := r.RunE(context.Background(),
		dash.TemplateSplitExpand("printf hi\n\n", ""),
		dash.WithStdoutFile(file1))
	if err != nil {
		panic(err)
	}

	err = r.RunE(context.Background(),
		dash.Split("cat"),
		dash.WithStdinFile(file1),
	)
	if err != nil {
		panic(err)
	}
	// Output: hi
}

func ExampleRunner_RunE() {
	r := dash.NewRunner()

	err := r.RunE(context.Background(),
		dash.TemplateSplitExpand(`ls {{.Dir}}`, struct{ Dir string }{
			Dir: "$HOME",
		}))

	if err != nil {
		panic(err)
	}
}
