package directcallertest

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"aduu.dev/utils/go/gomod"
)

const (
	currentPackage = "aduu.dev/utils/go/gomod/directcallertest"
)

func TestDirectCallerPackage(t *testing.T) {
	assert.Equal(t, currentPackage, gomod.DirectCallerPackagePanic(), "should get correct package")
}

func TestDirectCallerPackageE(t *testing.T) {
	gotPackage, err := gomod.DirectCallerPackageE()
	assert.NoError(t, err)
	assert.Equal(t, currentPackage, gotPackage, "should get correct package")
}

func TestDirectCallerFunctionName(t *testing.T) {
	assert.Equal(t, "TestDirectCallerFunctionName", gomod.DirectCallerFunctionName())
}

func TestDirectCallerFunctionNameWithPackageName(t *testing.T) {
	assert.Equal(t, "directcallertest.TestDirectCallerFunctionNameWithPackageName", gomod.DirectCallerFunctionNameWithPackageName())
}
