package main

import (
	"crypto/sha1"
	"fmt"
)

/**
	mencoba unutk menerapkan sha1 untuk enkripsi password
**/

func main() {
	var data = "password"
	var hash = sha1.New()
	hash.Write([]byte(data))
	var encripted = hash.Sum(nil)

	var encriptedString = fmt.Sprintf("%x", encripted)

	fmt.Println(encriptedString)
}
