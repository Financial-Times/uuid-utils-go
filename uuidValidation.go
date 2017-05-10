package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func validateUuid(contentUUID string) error {
	parsedUUID, err := uuid.FromString(contentUUID)
	if err != nil {
		return err
	}
	if contentUUID != parsedUUID.String() {
		return fmt.Errorf("Parsed UUID (%v) is different than the given uuid (%v).", parsedUUID, contentUUID)
	}
	return nil
}
