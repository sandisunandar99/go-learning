package main

import (
	"fmt"
	"os"
)

var path = "test.txt"

func CreateFile() {
	// cek dulu apakah di dalam folder sudah ada file atau belum
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			panic(err.Error())
		}
		defer file.Close()
	}

	fmt.Println("file telah di buat")
}

func main() {
	CreateFile()
}
