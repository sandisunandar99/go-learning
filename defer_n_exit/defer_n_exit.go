package main

import (
	"fmt"
	"os"
)

/**
	Defer digunakan untuk mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai.
	Exit digunakan untuk menghentikan program secara paksa (bener2 stack ga bisa dilanjutkan lagi)

**/

func main() {
	// contohDefer()
	contohExit()
}

// func contohDefer() {
// 	defer fmt.Println("awal")
// 	fmt.Println("akhir")
// }

func contohExit() {
	fmt.Println("awal")
	os.Exit(1) // dia akan behenti disini
	fmt.Println("akhir")
}
