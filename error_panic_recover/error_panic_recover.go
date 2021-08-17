package main

import (
	"errors"
	"fmt"
	"strings"
)

// contoh penggunaan error
func main() {
	defer catch()

	var input string
	fmt.Printf("ketikan angka : ")
	fmt.Scanln(&input)

	// var number int
	// var err error
	// number, err = strconv.Atoi(input)

	if valid, err := validate(input); valid {
		// fmt.Println(number, "adalah nomer")
		fmt.Println("cek : ", input)
	} else {
		// fmt.Println(input, "bukan nomor")
		// fmt.Println(err.Error())

		//menggunakan panic
		panic(err.Error())
	}

}

// menggunakan custom error
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("tidak bisa di kosongkan")
	}
	return true, nil
}

// menggunakan recover
func catch() {
	if err := recover(); err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("Tidak ada error")
	}
}
