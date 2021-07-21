package main

import (
	"fmt"
	"reflect"
)

/**
Reflection adalah teknik untuk inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkan memanipulasinya.
*/

func main() {

	var number = 32
	var valueOf = reflect.ValueOf(number)

	var nama = "Ujang dadang sagalana"
	var slice = reflect.ValueOf(nama)

	fmt.Println("Type : ", valueOf.Type())
	fmt.Println("Slice : ", slice.Slice(0, 5))
}
