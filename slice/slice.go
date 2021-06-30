package main

import "fmt"

func main() {
	var fruitsA = []string{"apple", "grape", "aple", "melon", "jambu"} // slice
	var fruitsB = [2]string{"banana", "melon"}                         // array
	var fruitsC = [...]string{"papaya", "grape"}                       // array

	fmt.Println(fruitsA)
	fmt.Println(fruitsB)
	fmt.Println(fruitsC)
	fmt.Println(fruitsA[1:4])

	aa := fruitsA[1:4]

	// fungsi len untuk menghitung lebar data pada array
	fmt.Println(len(fruitsA))

	//sama
	fmt.Println(cap(aa))
}
