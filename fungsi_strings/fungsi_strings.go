package main

import (
	"fmt"
	"strings"
)

func main() {
	//cek apakah string ada karakter yang sama
	var isExist = strings.Contains("Hello World", "World")
	fmt.Println(isExist)

	//unutk mendeteksi sebuat string yg diawali string tertentu
	var isPrefix = strings.HasPrefix("Hello World", "Hello")
	fmt.Println(isPrefix)

	//untuk mendeteksi string yang diakhiri string tertentu
	var isSuffix = strings.HasSuffix("Hello World", "World")
	fmt.Println(isSuffix)

	//count string huruf l
	var count = strings.Count("Hello World", "l")
	fmt.Println(count)

	// unutk mencari string dalam index ke berapa
	var index = strings.Index("Hello World", "World")
	fmt.Println(index)

	// replace string
	var replaced = strings.Replace("Hello World", "World", "Fungsi", 1)
	fmt.Println(replaced)

	// string repeat
	var repeated = strings.Repeat("Hello", 3)
	fmt.Println(repeated)

	// string split
	var split = strings.Split("Hello World ", " ")
	fmt.Println(split)

	// string join
	var jo = []string{"Hello", "World"}
	var joined = strings.Join(jo, ",")
	fmt.Println(joined)

	// tolower dan toUpper
	var tolower = strings.ToLower("Hello World")
	var toupper = strings.ToUpper("Hello World")

	fmt.Println(tolower)
	fmt.Println(toupper)

}
