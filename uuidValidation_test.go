package uuidutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ValidateUUID_EmptyString(t *testing.T) {
	err := ValidateUUID("")
	assert.NotNil(t, err, "Empty string was valid.")
}

func Test_ValidateUUID_FreeText(t *testing.T) {
	err := ValidateUUID("some random text")
	assert.NotNil(t, err, "Free text was valid.")
}

func Test_ValidateUUID_NoHyphens(t *testing.T) {
	validUUID := "0000ea7900a5504ea28d11bd108b35ac"
	err := ValidateUUID(validUUID)
	assert.NotNil(t, err, "UUID with no hyphens was valid.")
}

func Test_ValidateUUID_ValidUUIDLowercase(t *testing.T) {
	validUUID := "0000ea79-00a5-504e-a28d-11bd108b35ac"
	err := ValidateUUID(validUUID)
	assert.Nil(t, err, "Valid lowercase UUID got error on validation.")
}

func Test_ValidateUUID_ValidUUIDUppercase(t *testing.T) {
	validUUID := "0000EA79-00A5-504E-A28D-11BD108B35AC"
	err := ValidateUUID(validUUID)
	assert.Nil(t, err, "Valid uppercase UUID got error on validation.")
}

func Test_ValidateUUID_SmallerLength(t *testing.T) {
	validUUID := "0000ea79-00a5-504e-a28d-11bd108b35"
	err := ValidateUUID(validUUID)
	assert.NotNil(t, err, "UUID with smaller length was valid.")
}

func Test_ValidateUUID_BiggerLength(t *testing.T) {
	validUUID := "0000ea79-00a5-504e-a28d-11bd108b35ac11"
	err := ValidateUUID(validUUID)
	assert.NotNil(t, err, "UUID with bigger length was valid.")
}

func Test_ValidateUUID_InvalidCharacters(t *testing.T) {
	validUUID := "0000ea79-00a5-504e-a28d-11bd108b3XYZ"
	err := ValidateUUID(validUUID)
	assert.NotNil(t, err, "UUID with invalid characters was valid.")
}

func Test_ValidateUUID_MissingInnerGroup(t *testing.T) {
	validUUID := "0000ea79-00a5-a28d-11bd108b35ac"
	err := ValidateUUID(validUUID)
	assert.NotNil(t, err, "UUID with missing inner group was valid.")
}
