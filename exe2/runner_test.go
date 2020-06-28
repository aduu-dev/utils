package exe2

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"aduu.dev/utils/helper"
	"aduu.dev/utils/helper/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestRunner_MethodsReturnErrorIfSplitFailed(t *testing.T) {
	r := NewRunner()

	type testfunc func() error

	args := struct {
		ctx         context.Context
		splitResult SplitResult
	}{
		ctx:         context.Background(),
		splitResult: SplitResult{Err: errors.New("some-error")},
	}

	tests := []struct {
		name string
		fn   testfunc
	}{
		{
			name: "RunE",
			fn: func() error {
				return r.RunE(args.ctx, args.splitResult)
			},
		},
		{
			name: "RunWithOutputE",
			fn: func() error {
				_, err := r.RunWithOutputE(args.ctx, args.splitResult)
				return err
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.fn()

			if !assert.EqualError(t, err, args.splitResult.Err.Error()) {
				t.Fail()
			}
		})
	}
}

func TestRunner_UsesWithDir(t *testing.T) {
	tempDir := testhelper.MakeTempDir(t, "execute_UseWithDir")
	defer testhelper.DeleteTempDir(t, tempDir)

	// Creating a so i can try to touch a/b and so won't create a if it did not switch directories correctly
	a := filepath.Join(tempDir, "a")
	b := filepath.Join(tempDir, "a", "b")

	if err := os.Mkdir(a, 0777); err != nil {
		t.Fatalf("failed to create test dir a: %v", err)
	}

	r := NewRunner()

	err := r.RunE(context.Background(),
		TemplateSplitExpand(`touch a/b`, ""), WithDir(tempDir))

	if !assert.NoError(t, err) {
		t.Fail()
	}

	if !helper.DoesPathExist(b) {
		t.Errorf("failed to change directory into %s", tempDir)
		t.Fail()
	}

	t.Log("b does exist:", b)
}

func TestRunner_WithTimeout(t *testing.T) {
	r := NewRunner()

	err := r.RunE(context.Background(),
		TemplateSplitExpand(`sleep 1`, ""), WithTimeout(time.Millisecond*10))

	if !assert.EqualError(t, err, "signal: killed") {
		t.Fail()
	}

	err = r.RunE(context.Background(),
		TemplateSplitExpand(`sleep 0.01`, ""), WithTimeout(time.Second))

	if !assert.NoError(t, err) {
		t.Fail()
	}

}
