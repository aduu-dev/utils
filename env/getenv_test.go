package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"

	"aduu.dev/utils/go/gomod"
)

func randomEnvVariableName() string {
	return gomod.DirectCallerPackage() + "-" + uuid.New().String()
}

func TestGetenvRequired(t *testing.T) {
	key := randomEnvVariableName()

	assert.Panics(t, func() {
		GetenvRequired(key)
	}, "GetenvRequired should panic if not-set env variable is queried")

	wantValue := uuid.New().String()
	assert.NoError(t, os.Setenv(key, wantValue))

	value, err := GetenvE(key)
	assert.NoError(t, err)
	assert.Equal(t, wantValue, value)
}

func TestGetenvE(t *testing.T) {
	key := randomEnvVariableName()

	value, err := GetenvE(key)
	assert.Error(t, err)
	assert.Equal(t, "", value)

	wantValue := uuid.New().String()
	assert.NoError(t, os.Setenv(key, wantValue))

	value, err = GetenvE(key)
	assert.NoError(t, err)
	assert.Equal(t, wantValue, value)
}
