package uuidutils

import (
	"crypto/sha1"
	"encoding/binary"
)

const (
	NAMESPACE_URL_ID = "6ba7b811-9dad-11d1-80b4-00c04fd430c8"
)

// NewGenerateUUID gives a new V5 UUID based on a random string
// Corresponds to com.ft.uuidutils.GenerateUuid#from(final String string)
func NewV5UUIDFrom(str string) UUID {
	digest := newV5Digest(NAMESPACE_URL_ID, str)

	hi := createMSB(digest)
	lo := createLSB(digest)

	return UUID{binary.BigEndian.Uint64(hi), binary.BigEndian.Uint64(lo)}
}

func newV5Digest(ns string, str string) []byte {
	h := sha1.New()
	h.Write([]byte(ns))
	h.Write([]byte(str))

	return h.Sum(nil)
}

func createMSB(digest []byte) []byte {
	hi := make([]byte, 8)
	hi[2] = digest[0]
	hi[3] = digest[1]
	hi[5] = digest[2]
	hi[7] = digest[3]
	hi[6] |= 5 << 4
	return hi
}

func createLSB(digest []byte) []byte {
	lo := make([]byte, 8)
	lo[0] = digest[8]
	lo[0] &= 0xbf
	lo[0] |= 0x80
	lo[1] = digest[9]
	copy(lo[2:8], digest[10:16])
	return lo
}
