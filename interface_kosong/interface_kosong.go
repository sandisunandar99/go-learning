package main

import "fmt"

/**
interface kosong adalah suatu variable spesial dimana bisa menampung tipe data apapun
seperti array, integer, string, bool, dsb
*/

func main() {
	var random interface{}

	random = "sonando sota sote"
	fmt.Println(random)

	random = []string{"satu", "dua", "tiga"}
	fmt.Println(random)

	var siswa = []map[string]interface{}{
		{"nama": "soni", "kelas": 1},
		{"nama": "dodo", "kelas": 2},
	}

	for _, val := range siswa {
		fmt.Println("nama :", val["nama"], "kelas :", val["kelas"])
	}

}
