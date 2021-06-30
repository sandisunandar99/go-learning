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

	// looping menggunakan range
	num := []int{5, 2, 10, 4, 11, 61, 87, 8, 90}

	for key, val := range num {
		fmt.Printf("key %d ", key)
		fmt.Printf("val %d ", val)
	}
}
