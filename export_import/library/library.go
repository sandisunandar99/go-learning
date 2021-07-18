package library

import "fmt"

//unutk tipe struck jika public maka variable didalamnya harus capital juga seperti nama maka huruf 'N' nya harus capital
type Student struct {
	Nama  string
	Nilai int
}

// ini adalah fungsi public dan dapat di cirikan karakter awal Hello huruf awal nya capital "H"
func HelloWorld() {
	fmt.Println("ini adalah fungsi public")
	hello()
}

// dan ini fungsi privare dengan tanda hurup awal nya kecil "h"
func hello() {
	fmt.Println("ini adalah fungsi private")
}
