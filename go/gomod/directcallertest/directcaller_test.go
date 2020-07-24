package directcallertest

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"aduu.dev/utils/go/gomod"
)

const (
	currentPackage = "aduu.dev/utils/go/gomod/directcallertest"
)

func runningWithBazel(t *testing.T) bool {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if strings.Contains(wd, "bazel-out") {
		t.Skip()
		return true
	}

	return false
}

func TestDirectCallerPackage(t *testing.T) {
	if runningWithBazel(t) {
		return
	}

	assert.Equal(t, currentPackage, gomod.DirectCallerPackagePanic(), "should get correct package")
}

func TestDirectCallerPackageE(t *testing.T) {
	if runningWithBazel(t) {
		return
	}

	gotPackage, err := gomod.DirectCallerPackageE()
	assert.NoError(t, err)
	assert.Equal(t, currentPackage, gotPackage, "should get correct package")
}

func TestDirectCallerFunctionName(t *testing.T) {
	if runningWithBazel(t) {
		return
	}

	assert.Equal(t, "TestDirectCallerFunctionName", gomod.DirectCallerFunctionName())
}

func TestDirectCallerFunctionNameWithPackageName(t *testing.T) {
	if runningWithBazel(t) {
		return
	}

	assert.Equal(t, "directcallertest.TestDirectCallerFunctionNameWithPackageName", gomod.DirectCallerFunctionNameWithPackageName())
}
