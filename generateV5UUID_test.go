package uuidutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"net/url"
)

func Test_NewV5UUIDFrom_IsSuccessful(t *testing.T) {
	urlRadu, _ := url.Parse("http://api.ft.com/duders/Radu")
	urlBogdan, _ := url.Parse("http://api.ft.com/duders/Bogdan")
	urlHuni, _ := url.Parse("http://api.ft.com/duders/Huni")
	urlGeorge, _ := url.Parse("http://api.ft.com/duders/George")
	urlRandomID, _ := url.Parse("http://api.ft.com/things/U11603507121721xBE")

	resultUUIDRadu := NewV5UUIDFromURL(urlRadu)
	resultUUIDBogdan := NewV5UUIDFromURL(urlBogdan)
	resultUUIDHuni := NewV5UUIDFromURL(urlHuni)
	resultUUIDGeorge := NewV5UUIDFromURL(urlGeorge)
	resultUUIDRandomID := NewV5UUIDFromURL(urlRandomID)

	assert.Equal(t, "0000ea79-00a5-504e-a28d-11bd108b35ac", resultUUIDRadu.String(), "Expected UUID didn't matched")
	assert.Equal(t, "000025f3-00da-50bc-a6b1-271b74f9130f", resultUUIDBogdan.String(), "Expected UUID didn't matched")
	assert.Equal(t, "0000f4e2-002c-509b-9a1b-54ec6e055973", resultUUIDHuni.String(), "Expected UUID didn't matched")
	assert.Equal(t, "000097a3-0033-509e-8ad9-a9ac5257a7cb", resultUUIDGeorge.String(), "Expected UUID didn't matched")
	assert.Equal(t, "00006a43-00dd-5084-ba5e-1b05d24a06e0", resultUUIDRandomID.String(), "Expected UUID didn't matched")
}

func Test_NewV5UUIDFromURL_IsRepeatable(t *testing.T) {
	urlBogdan, _ := url.Parse("http://api.ft.com/duders/Radu")

	firstGeneratedUUID := NewV5UUIDFromURL(urlBogdan)
	secondGeneratedUUID := NewV5UUIDFromURL(urlBogdan)

	assert.Equal(t, firstGeneratedUUID.String(), secondGeneratedUUID.String(), "Derived original UUID didn't matched same derived UUID.")
}
