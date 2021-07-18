package main

import "fmt"

// fungsi variadic biasanya ditandai dengan (...)
func main() {
	hasil := calc(1, 1, 1, 1, 1, 1, 1)

	fmt.Printf("hasil =  %.0f \n", hasil)

	ikan(3, "juaer", "gurame", "tomang")
}

// tanda variadic ialah pada fungsi terdapat multiple variabel dengan tipe data sama (...)
func calc(n ...int) float64 {
	var sum int = 0

	for _, num := range n {
		sum += num
	}

	return float64(sum)
}

func ikan(jumlah int, ikan ...string) {
	fmt.Println(jumlah)
	fmt.Println(ikan)
}
