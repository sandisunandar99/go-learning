package main

import "fmt"

func main() {

	var name [2]string
	name[0] = "satu"
	name[1] = "dua"

	fmt.Println(name)

	var buah = [3]string{"mangga", "apel", "pisang"}

	fmt.Println(buah)

	var binatang = [...]string{"jaguar", "kucing", "ayam", "entog"}

	fmt.Println(binatang)
	fmt.Println(len(binatang)) // unutk menghitung lebar array

	//loping array
	for _, val := range binatang {
		fmt.Println(val)
	}

	// alokasi array menggunakan keyword
	var burung = make([]string, 2)
	burung[0] = "piit"
	burung[1] = "kakatua"

	fmt.Println(burung)
}
