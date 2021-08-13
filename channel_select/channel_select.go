package main

import (
	"fmt"
	"runtime"
)

/**
fungsi utama channel bukan untuk kontrol, melainkan untuk sharing data antar goroutine.
*/

func main() {
	runtime.GOMAXPROCS(2)

	var number = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Number : ", number)

	var ch1 = make(chan float64)
	go getAvg(number, ch1)

	var ch2 = make(chan int)
	go getMax(number, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Println("Avg : ", avg)
		case max := <-ch2:
			fmt.Println("Max : ", max)
		}
	}

}

func getAvg(num []int, ch chan float64) {
	var sum = 0
	for _, v := range num {
		sum += v
	}

	ch <- float64(sum) / float64(len(num))
}

func getMax(num []int, ch chan int) {
	var max = num[0]

	for _, v := range num {
		if max < v {
			max = v
		}
	}

	ch <- max
}
