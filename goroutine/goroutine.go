package main

import (
	"fmt"
	"runtime"
)

/**
goroutine disebut juga mini thread
Eksekusi goroutine bersifat asynchronous, menjadikannya tidak saling tunggu dengan goroutine lain.
*/

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	// untuk menentukan jumlah core yang di aktifkan pada eksekusi program dibawahnya
	runtime.GOMAXPROCS(2)

	// bagian go routine yang ditandai dengan "go" seperti dibawah ini
	go print(5, "hello")
	print(5, "apa kabar ")

	var input string
	fmt.Scanln(&input)
}
