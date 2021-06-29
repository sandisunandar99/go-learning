package main

import "fmt"

func main() {
	//perulangan menggunakan for
	var count = 3

	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	var k = 0
	for k < 5 {
		fmt.Println(k)
		k++
	}
}
