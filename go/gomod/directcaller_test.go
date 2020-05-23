package gomod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectCallerPackage(t *testing.T) {
	wantPackage := "aduu.dev/pkg/go/gomod"
	assert.Equal(t, DirectCallerPackage(), wantPackage, "should get correct package")
}

func TestDirectCallerFunctionName(t *testing.T) {
	assert.Equal(t, "TestDirectCallerFunctionName", DirectCallerFunctionName())
}

func TestDirectCallerFunctionNameWithPackageName(t *testing.T) {
	assert.Equal(t, "gomod.TestDirectCallerFunctionNameWithPackageName", DirectCallerFunctionNameWithPackageName())
}
