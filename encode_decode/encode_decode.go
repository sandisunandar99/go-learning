package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "santos delroso"

	var encodeString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encode :", encodeString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodeString)
	var decodeString = string(decodeByte)
	fmt.Println("decode :", decodeString)
}
