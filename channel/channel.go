package main

import (
	"fmt"
	"runtime"
)

/**
Channel digunakan untuk menghubungkan goroutine satu dengan goroutine lain.
Mekanisme yang dilakukan adalah serah-terima data lewat channel tersebut.
Dalam komunikasinya, sebuah channel difungsikan sebagai pengirim di sebuah goroutine,
dan juga sebagai penerima di goroutine lainnya.
Pengiriman dan penerimaan data pada channel bersifat blocking atau synchronous.
*/

func main() {
	runtime.GOMAXPROCS(2)

	var msg = make(chan string)

	var Hallo = func(w string) {
		var data = fmt.Sprintf("hello %s", w)
		msg <- data
	}

	go Hallo("sandi")
	go Hallo("ujang")
	go Hallo("babang")

	var msgBackup1 = <-msg

	fmt.Println(msgBackup1)

	var msgBackup2 = <-msg

	fmt.Println(msgBackup2)

	var msgBackup3 = <-msg

	fmt.Println(msgBackup3)

}
