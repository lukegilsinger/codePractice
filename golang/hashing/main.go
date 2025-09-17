package main

// go run .
// go run main.go

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {

	var SourceSystemCode interface{}
	var DigitalIdentityId interface{}

	SourceSystemCode = "US3"
	DigitalIdentityId = "12345"
	attributesMap := map[string]any{
		"SourceSystemCode":  SourceSystemCode,
		"DigitalIdentityId": DigitalIdentityId,
	}

	primaryKeysMd5 := md5.Sum([]byte(fmt.Sprintf("<%s\u03A3%s>", attributesMap["SourceSystemCode"], attributesMap["DigitalIdentityId"])))
	id := hex.EncodeToString(primaryKeysMd5[:])

	fmt.Println("Unique ID: ", id)
}
