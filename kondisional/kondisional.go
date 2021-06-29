package main

import "fmt"

func main() {
	point := 4

	// kondisional if else
	if point == 10 {
		fmt.Println("di skors")
	} else if point > 5 {
		fmt.Println("sp3")
	} else if point == 4 {
		fmt.Println("sp2")
	} else {
		fmt.Println("sp1")
	}

	// kondisi menggunakan switch
	switch point {
	case 8, 9, 10:
		fmt.Println("sempurna")
	case 7, 6:
		fmt.Println("lumayan")
	default:
		fmt.Println("tidak lulus")
	}

}
