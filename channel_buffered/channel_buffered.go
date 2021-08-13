package main

import (
	"fmt"
	"runtime"
)

/**
Proses transfer data pada channel secara default dilakukan dengan cara un-buffered,
tidak di-buffer di memori. Ketika terjadi proses kirim data via channel dari sebuah goroutine,
maka harus ada goroutine lain yang bertugas menerima data dari channel yang sama,
dengan proses serah-terima yang bersifat blocking.
Maksudnya, baris kode setelah kode pengiriman dan penerimaan data tidak akan diproses sebelum proses serah-terima itu sendiri selesai.
*/

func main() {
	runtime.GOMAXPROCS(2)

	var msg = make(chan int, 2)

	// buffer data dulu disini dan setelah itu baru di keluarkan hasilnya.
	go func() {
		for {
			i := <-msg
			fmt.Printf("data %d ", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data ", i)
		msg <- i
	}
}
