package uuidutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewSingleDigestedV3UUID_IsSuccessful(t *testing.T) {
	resultUUID := NewV3UUID("some random text")

	assert.Equal(t, "07671a03-8c0e-3437-a3d4-21693b073c3b", resultUUID.String())
}

func Test_NewSingleDigestedV3UUID_IsRepeatable(t *testing.T) {
	randomText := "some random text"

	firstGeneratedUUID := NewV3UUID(randomText)
	secondGeneratedUUID := NewV3UUID(randomText)

	assert.Equal(t, firstGeneratedUUID.String(), secondGeneratedUUID.String(), "Derived original UUID didn't matched same derived UUID.")
}

func Test_NewDoubleDigestedV3UUID_IsSuccessful(t *testing.T) {
	resultUUID := NewDoubleDigestedV3UUID("some random text")

	assert.Equal(t, "5cdd4e45-5f40-3e8f-a75a-29e051ae6b37", resultUUID.String())
}

func Test_NewDoubleDigestedV3UUID_IsRepeatable(t *testing.T) {
	randomText := "some random text"

	firstGeneratedUUID := NewDoubleDigestedV3UUID(randomText)
	secondGeneratedUUID := NewDoubleDigestedV3UUID(randomText)

	assert.Equal(t, firstGeneratedUUID.String(), secondGeneratedUUID.String(), "Derived original UUID didn't matched same derived UUID.")
}
