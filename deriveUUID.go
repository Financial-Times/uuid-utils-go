package uuidutils

import (
	"strconv"
	"strings"

	"github.com/willf/bitset"
)

type Salt string

const (
	IMAGE_SET Salt = "imageset"
)

type deriveUUID struct {
	saltUUIDLsb *bitset.BitSet
}

// NewUUIDDeriverWith creates a UUID deriver with the given salt.
func NewUUIDDeriverWith(salt Salt) *deriveUUID {
	return &deriveUUID{saltUUIDLsb: saltToUUIDLsb(salt)}
}

// From derives a new UUID based on another UUID using the internal salt.
// The function is reversible, so using it again on the derived UUID will obtain the original UUID.
func (du *deriveUUID) From(originalUUID *UUID) (*UUID, error) {
	return otherUUID(originalUUID, du.saltUUIDLsb)
}

func otherUUID(otherUUID *UUID, saltUUIDLsb *bitset.BitSet) (*UUID, error) {
	uuidBits := toBitSet(otherUUID.lsb)
	uuidBits.InPlaceSymmetricDifference(saltUUIDLsb) //XOR

	lsb, err := strconv.ParseUint(strings.TrimSuffix(uuidBits.DumpAsBits(), "."), 2, 64)
	return &UUID{otherUUID.msb, uint64(lsb)}, err
}

func saltToUUIDLsb(salt Salt) *bitset.BitSet {
	return toBitSet(NewNameUUIDFromBytes([]byte(string(salt))).lsb)
}

func toBitSet(number uint64) *bitset.BitSet {
	lsbString := strconv.FormatUint(number, 2)

	var b bitset.BitSet
	for i, character := range lsbString {
		if character == '1' {
			b.Set(uint(len(lsbString)-1) - uint(i))
		}
	}
	return &b
}
