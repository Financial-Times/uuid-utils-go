package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_ValidateUUID_ValidUUIDLowercase(t *testing.T) {
	validUUID := "0000ea79-00a5-504e-a28d-11bd108b35ac"
	err := ValidateUUID(validUUID)
	assert.Nil(t, err, "Valid UUID got error on validation.")
}

func Test_ValidateUUID_ValidUUIDUppercase(t *testing.T) {
	validUUID := "0000EA79-00A5-504E-A28D-11BD108B35AC"
	err := ValidateUUID(validUUID)
	assert.Nil(t, err, "Valid UUID got error on validation.")
}

func Test_ValidateUUID_EmptyStringUUID(t *testing.T) {
	err := ValidateUUID("")
	assert.NotNil(t, err, "Valid UUID got error on validation.")
}

func Test_ValidateUUID_SmallerLength(t *testing.T) {
	validUUID := "0000ea79-00a5-504e-a28d-11bd108b35"
	err := ValidateUUID(validUUID)
	assert.NotNil(t, err, "Valid UUID got error on validation.")
}

func Test_ValidateUUID_BiggerLength(t *testing.T) {
	validUUID := "0000ea79-00a5-504e-a28d-11bd108b35ac11"
	err := ValidateUUID(validUUID)
	assert.NotNil(t, err, "Valid UUID got error on validation.")
}

func Test_ValidateUUID_InvalidCharacters(t *testing.T) {
	validUUID := "0000ea79-00a5-504e-a28d-11bd108b3XYZ"
	err := ValidateUUID(validUUID)
	assert.NotNil(t, err, "Valid UUID got error on validation.")
}

func Test_ValidateUUID_MissingInnerGroup(t *testing.T) {
	validUUID := "0000ea79-00a5-a28d-11bd108b35ac"
	err := ValidateUUID(validUUID)
	assert.NotNil(t, err, "Valid UUID got error on validation.")
}
