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

func ExampleRunner_withTimeout() {
	r := exe2.NewRunner()

	err := r.RunE(context.Background(),
		exe2.TemplateSplitExpand(`sleep 1`, ""),
		exe2.WithTimeout(time.Millisecond*10))

	if err != nil {
		panic(err)
	}
}

func ExampleRunner_withStdPipes() {
	file1 := filepath.Join(os.TempDir(), "file1")
	//file2 := filepath.Join(os.TempDir(), "file2")

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
