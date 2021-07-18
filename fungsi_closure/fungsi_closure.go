package main

import "fmt"

// fungsi yang disimpan dalam variable pada suatu fungsi tertentu
func main() {
	var minMax = func(n []int) (int, int) {
		var min, max int

		for key, val := range n {

			if key == 0 {
				min, max = val, val
			}

			if val < min {
				min = val
			}

			if val > max {
				max = val
			}
		}

		return min, max
	}

	var number = []int{10, 22, 3, 77, 1, 5}

	var min, max = minMax(number)

	fmt.Printf("min : %d ", min)
	fmt.Printf("max : %d", max)
}
