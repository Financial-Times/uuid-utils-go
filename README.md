# uuid-utils-go
[![Circle CI](https://circleci.com/gh/Financial-Times/uuid-utils-go/tree/master.png?style=shield)](https://circleci.com/gh/Financial-Times/uuid-utils-go/tree/master)[![Go Report Card](https://goreportcard.com/badge/github.com/Financial-Times/uuid-utils-go)](https://goreportcard.com/report/github.com/Financial-Times/uuid-utils-go) [![Coverage Status](https://coveralls.io/repos/github/Financial-Times/uuid-utils-go/badge.svg)](https://coveralls.io/github/Financial-Times/uuid-utils-go)

## Introduction

This ia a library containing UUID generation and validation used in UPP's Golang services.

#### What methods should you use?

1. Generating an UUID based on another UUID: methods in deriveUUID.go

2. Generating an UUID from an random string (V3 version, i.e. using MD5 hashing): methods in generateV3UUID.go

3. Generating an UUID from an URL (V5 version, i.e. using SHA1 hashing): methods in generateV5UUID.go

4. Validating an UUID: methods in uuidValidation.go

Note: this library also contains methods that replicate java.util.UUID class (in uuid.go), for cases where it's already in use.

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

1. Generating an UUID based on another UUID:

            originalUUID, _ := uuidutils.NewUUIDFromString("0000ea79-00a5-504e-a28d-11bd108b35ac")
            uuidDeriver := uuidutils.NewUUIDDeriverWith(uuidutils.IMAGE_SET)

            derivedUUID := uuidDeriver.From(originalUUID)
            
            revertToOriginalUUID := uuidDeriver.From(derivedUUID)
            
2. Generating an UUID from an random string (V3 version, i.e. using MD5 hashing):

            someString := "1AB23ad1x34"
            generatedUUID := uuidutils.NewV3UUID(someString)

3. Generating an UUID from an URL (V5 version, i.e. using SHA1 hashing):

            someString := "1AB23ad1x34"
            generatedUUID := uuidutils.NewDoubleDigestedV3UUID(someString)

4. Validating an UUID:

            someUUID := "0000ea79-00a5-504e-a28d-11bd108b35ac"
            err := ValidateUUID(validUUID)
            if err != nil {
                // do something with the error
            }