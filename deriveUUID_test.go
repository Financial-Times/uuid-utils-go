package uuidutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var uuidDeriver = NewUUIDDeriverWith(IMAGE_SET)
var otherDeriver = NewUUIDDeriverWith(Salt("other salt"))

func Test_From_IsSuccessful(t *testing.T) {
	originalUUID1, _ := NewUUIDFromString("0000ea79-00a5-504e-a28d-11bd108b35ac")
	originalUUID2, _ := NewUUIDFromString("000025f3-00da-50bc-a6b1-271b74f9130f")
	originalUUID3, _ := NewUUIDFromString("0000f4e2-002c-509b-9a1b-54ec6e055973")
	originalUUID4, _ := NewUUIDFromString("000097a3-0033-509e-8ad9-a9ac5257a7cb")
	originalUUID5, _ := NewUUIDFromString("00006a43-00dd-5084-ba5e-1b05d24a06e0")

	derivedUUID1, _ := uuidDeriver.From(originalUUID1)
	derivedUUID2, _ := uuidDeriver.From(originalUUID2)
	derivedUUID3, _ := uuidDeriver.From(originalUUID3)
	derivedUUID4, _ := uuidDeriver.From(originalUUID4)
	derivedUUID5, _ := uuidDeriver.From(originalUUID5)

	assert.Equal(t, "0000ea79-00a5-504e-3ceb-8627caff9ee5", derivedUUID1.String(), "Expected UUID didn't matched")
	assert.Equal(t, "000025f3-00da-50bc-38d7-b081ae8db846", derivedUUID2.String(), "Expected UUID didn't matched")
	assert.Equal(t, "0000f4e2-002c-509b-047d-c376b471f23a", derivedUUID3.String(), "Expected UUID didn't matched")
	assert.Equal(t, "000097a3-0033-509e-14bf-3e3688230c82", derivedUUID4.String(), "Expected UUID didn't matched")
	assert.Equal(t, "00006a43-00dd-5084-2438-8c9f083eada9", derivedUUID5.String(), "Expected UUID didn't matched")
}

func Test_From_SameSaltSameResults(t *testing.T) {
	originalUUID, _ := NewUUIDFromString("0000ea79-00a5-504e-a28d-11bd108b35ac")
	sameUUID, _ := NewUUIDFromString("0000ea79-00a5-504e-a28d-11bd108b35ac")

	derivedOriginalUUID, _ := uuidDeriver.From(originalUUID)
	derivedSameUUID, _ := uuidDeriver.From(sameUUID)

	assert.Equal(t, derivedOriginalUUID.String(), derivedSameUUID.String(), "Derived original UUID didn't matched same derived UUID.")
}

func Test_From_DifferentSaltsDifferentResults(t *testing.T) {
	someUUID, _ := NewUUIDFromString("0000ea79-00a5-504e-a28d-11bd108b35ac")
	firstSaltDerivedUUID, _ := uuidDeriver.From(someUUID)
	secondSaltDerivedUUID, _ := otherDeriver.From(someUUID)

	assert.NotEqual(t, firstSaltDerivedUUID.String(), secondSaltDerivedUUID.String(), "Derived original UUID didn't matched same derived UUID.")
}

func Test_From_UUIDIsReversible(t *testing.T) {
	originalUUID, _ := NewUUIDFromString("0000ea79-00a5-504e-a28d-11bd108b35ac")

	derivedOriginalUUID, _ := uuidDeriver.From(originalUUID)
	reversedOriginalUUID, _ := uuidDeriver.From(derivedOriginalUUID)

	assert.Equal(t, originalUUID.String(), reversedOriginalUUID.String(), "Derived original UUID didn't matched same derived UUID.")
}
