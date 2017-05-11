# uuid-utils-go
[![Circle CI](https://circleci.com/gh/Financial-Times/uuid-utils-go/tree/master.png?style=shield)](https://circleci.com/gh/Financial-Times/uuid-utils-go/tree/master)[![Go Report Card](https://goreportcard.com/badge/github.com/Financial-Times/uuid-utils-go)](https://goreportcard.com/report/github.com/Financial-Times/uuid-utils-go) [![Coverage Status](https://coveralls.io/repos/github/Financial-Times/uuid-utils-go/badge.svg)](https://coveralls.io/github/Financial-Times/uuid-utils-go)

## Introduction

This ia a library containing UUID generation and validation used in UPP's Golang services.

There are 3 categories of operations offered (check examples at the bottom):
- deriving an UUID from another UUID (and reverting): in deriveUUID.go
- generating an UUID from a random string: in generateUUID.go
- validating an UUID: in uuidValidation.go

Note: this library also contains methods that replicate java.util.UUID class (in uuid.go), which also offers a method to generate UUID (i.e. NewNameUUIDFromBytes) that uses the old V3 UUID implementation, instead of V5, so it is advised not to use it if possible (use the method NewV5UUIDFrom).

## Installation:

1. Import the library (in your code):

        import (
                uuidutils "github.com/Financial-Times/uuid-utils-go"
        )

2. Vendor it (e.g. with govendor):

        govendor fetch github.com/Financial-Times/uuid-utils-go@1.0.0
        
Note1: this assumes your repo is already vendored (i.e. you ran: govendor init)

Note2: change in the fetch example above the appropriate tag release you want: @x.y.z

## Usage examples:

1. Deriving an UUID from another UUID (and reverting):

            originalUUID, _ := uuidutils.NewUUIDFromString("0000ea79-00a5-504e-a28d-11bd108b35ac")
            uuidDeriver := uuidutils.NewUUIDDeriverWith(uuidutils.IMAGE_SET)

            derivedUUID := uuidDeriver.From(originalUUID)
            
            revertToOriginalUUID := uuidDeriver.Revert(derivedUUID)
            
2. Generate an UUID from a random string:

            someString := "1AB23ad1x34"
            generatedUUID := uuidutils.NewV5UUIDFrom(someString)

3. Validate an UUID:

            someUUID := "0000ea79-00a5-504e-a28d-11bd108b35ac"
            err := ValidateUUID(validUUID)
            if err != nil {
                // do something with the error
            }
