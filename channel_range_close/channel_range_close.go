package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var message = make(chan string)
	go kirimPesan(message)
	printMsg(message)
}

func kirimPesan(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("Data %d", i)
	}

}

func printMsg(ch <-chan string) {
	for data := range ch {
		fmt.Printf("%s\n", data)
	}
}
