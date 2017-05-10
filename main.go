package main

import (
	"fmt"
)

func main() {
	idParam := "http://api.ft.com/duders/Bogdan"

	uuidGenerator := NewGenerateUUID()
	resultUUID, _ := uuidGenerator.From(idParam)

	fmt.Println(resultUUID.String())
}