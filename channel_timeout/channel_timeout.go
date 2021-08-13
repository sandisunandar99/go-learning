package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

/***
	Teknik channel timeout digunakan untuk mengontrol penerimaan data dari channel mengacu ke kapan waktu diterimanya data,
	dengan durasi timeout bisa kita tentukan sendiri.
**/

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan int)
	go KirimData(ch)
	retriveData(ch)

}

func KirimData(ch chan int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retriveData(ch chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Println("Data diterima ", data)
		case <-time.After(time.Duration(5) * time.Second):
			fmt.Println("Timeout")
			break loop
		}
	}
}
