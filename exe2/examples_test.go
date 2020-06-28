package exe2_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"aduu.dev/utils/exe2"
)

func ExampleRunner_RunWithOutputE() {
	r := exe2.NewRunner()

	out, err := r.RunWithOutputE(
		context.Background(),
		exe2.TemplateSplitExpand(`echo hi`, ""))
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
	// Output: hi
}

func ExampleRunner_withDir() {
	r := exe2.NewRunner()

	err := r.RunE(context.Background(),
		exe2.TemplateSplitExpand(`ls`, ""),
		exe2.WithDir(os.Getenv("HOME")),
	)
	if err != nil {
		panic(err)
	}
}

func ExampleRunner_withDeadline() {
	r := exe2.NewRunner()

	_ = 3

	err := r.RunE(context.Background(),
		exe2.TemplateSplitExpand(exe2.SleepCommand("1"), ""),
		exe2.WithTimeout(time.Microsecond))

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

	r := exe2.NewRunner()

	err := r.RunE(context.Background(),
		exe2.TemplateSplitExpand("printf hi\n\n", ""),
		exe2.WithStdoutFile(file1))
	if err != nil {
		panic(err)
	}

	err = r.RunE(context.Background(),
		exe2.Split("cat"),
		exe2.WithStdinFile(file1),
	)
	if err != nil {
		panic(err)
	}
	// Output: hi
}

func ExampleRunner_RunE() {
	r := exe2.NewRunner()

	err := r.RunE(context.Background(),
		exe2.TemplateSplitExpand(`ls {{.Dir}}`, struct{ Dir string }{
			Dir: "$HOME",
		}))

	if err != nil {
		panic(err)
	}
}
