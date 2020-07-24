package packages

import (
	"context"
	"testing"

	"aduu.dev/utils/dash"
	"aduu.dev/utils/dash/dmock"
	"github.com/stretchr/testify/assert"
)

func TestListPackages(t *testing.T) {
	m := dmock.New()

	m.On("RunWithOutputE",
		context.Background(),
		&dash.SplitResult{
			Name: "go",
			Args: []string{
				"list", "./...",
			},
		},
		&dash.ExecuteSetting{Dir: "abc"},
	).Return(
		`my.test
my.test/a
my.test/b`,
		nil,
	)

	want := []string{
		"my.test",
		"my.test/a",
		"my.test/b",
	}
	got, err := ListPackages(m, "abc")
	assert.NoError(t, err)
	assert.Equal(t, want, got)

	m.AssertExpectations(t)
}
