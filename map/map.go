package main

import "fmt"

func main() {

	var burung = map[string]int{}

	burung["gagak"] = 10
	burung["jalak"] = 5

	fmt.Println(burung["jalak"])
	fmt.Println(burung)

	var tanggal = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range tanggal {
		fmt.Println("key:", key, "--- val :", val)
	}

	delete(tanggal, "april")

	fmt.Println(tanggal)

	// contoh random key value pada map
	var data = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	fmt.Println(data)
	fmt.Println(data[1])
	fmt.Println(data[1]["id"])
}
