package uuidutils

import (
	"fmt"
	"strings"
)

func ValidateUUID(contentUUID string) error {
	parsedUUID, err := NewUUIDFromString(contentUUID)
	if err != nil || strings.ToLower(contentUUID) != parsedUUID.String() {
		return fmt.Errorf("Given UUID: (%v) is invalid!", contentUUID)
	}
	return nil
}
