package uuidutils

import (
	"crypto/md5"
)

// NewV3UUID generates a new V3 UUID based on a random string
// Corresponds to com.ft.uuidutils.GenerateV3UUID#singleDigested(String data)
func NewV3UUID(data string) *UUID {
	return newNameUUIDFromBytes([]byte(data))
}

// NewDoubleDigestedV3UUID generates a new V3 UUID based on a random string
// Corresponds to com.ft.uuidutils.GenerateV3UUID#doubleDigested(String data)
func NewDoubleDigestedV3UUID(data string) *UUID {
	firstDigest := md5.Sum([]byte(data))
	return newNameUUIDFromBytes(firstDigest[:])
}
