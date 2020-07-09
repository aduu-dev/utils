package dash

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_TemplateSplitExpand_dont_panic_on_template_error(t *testing.T) {
	_ = TemplateSplitExpand(`echo ${{.nonexistentKey}}`, struct {
		name string
	}{
		name: fmt.Sprintf("$%s", ""),
	})
}

func Test_TemplateSplitExpand(t *testing.T) {
	key := uuid.New().String()

	t.Logf("Setting Env key %s", key)
	os.Setenv(key, "abc")

	t.Cleanup(func() {
		t.Log("unsetting env", "key", key)
		os.Unsetenv(key)
	})

	got := TemplateSplitExpand(`echo ${{.Name}}`, struct {
		Name string
	}{
		Name: fmt.Sprintf("$%s", key),
	})

	if !assert.NoError(t, got.Err) {
		t.Fail()
	}

	want := &SplitResult{Name: "echo", Args: []string{key}}

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}

func Test_SplitTemplateExpand(t *testing.T) {
	key := uuid.New().String()

	t.Logf("Setting Env key %s", key)
	os.Setenv(key, "abc")

	t.Cleanup(func() {
		t.Log("Unsetting env", "key", key)
		os.Unsetenv(key)
	})

	got := SplitTemplateExpand(`echo ${{.Name}}`, struct {
		Name string
	}{
		Name: fmt.Sprintf("$%s def", key),
	})

	if !assert.NoError(t, got.Err) {
		t.Fail()
	}

	want := &SplitResult{Name: "echo", Args: []string{key + " def"}}

	if !assert.Equal(t, want, got) {
		t.Fail()
	}
}
