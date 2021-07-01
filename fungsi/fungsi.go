package main

import "fmt"

func main() {

	//pemangilan fungsi tanpa return dari fungsi yang dipanggil
	OutString("hellow bro")

	name := []string{"satu", "dua", "tiga"}
	DataArray(name)

	//pemanggilan fungsi dengan return
	hasil := HitungAngka(5, 2)
	fmt.Println(hasil)
}

func OutString(msg string) {
	fmt.Println(msg)
}

func DataArray(arr []string) {
	fmt.Println(arr)
}

// ketika fungsi harus ada return, maka dalam fungsi tsb harus di definisikan type datanya atau dalam bentuk struct
func HitungAngka(a int, b int) int {
	var hasil = a + b

	return hasil
}
