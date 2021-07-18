package main

import (
	"fmt"

	"github.com/go-basic/export_import/library"
)

/*
	Export / Import module dapat diklasifikasikan menjadi 2
	1. module Public => yang ditandain dengan hurup capital pada awal karakter fungsinnya ex: HelloWorld()
	2. module private => yang ditandain dengan hudup kecil pada awal karakter fungsinya ex: helloWorld()

	catatan :
	fungsi yang hanya bisa di panggil antar module adalah fungsi public, dan untuk private hanya dapat dipanggil
	di module dirinya sendiri.
*/

func main() {
	library.HelloWorld()

	var s = library.Student{"sandi", 10}

	fmt.Println("Nama : ", s.Nama)
	fmt.Println("Nilai : ", s.Nilai)
}
