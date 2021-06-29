package main

import "fmt"

func main() {

	var name string = "sandi" // mendefinisikan variable dengan tipe data
	last := "sunandar"        // define variable tanpa definisi type data

	var number int = 6
	koma := 0.27

	var satu, dua, tiga string = "satu", "dua", "tiga" // definisi multiple variable

	num, _ := 1, 2 // underscore adalah variable yang bisa digunakan atau pun tidak

	fmt.Printf("hallo %s %s \n", name, last)
	fmt.Printf("number %d %.2f \n", number, koma)
	fmt.Printf("number %s , %s , %s \n", satu, dua, tiga)
	fmt.Printf("number %d", num)
}
