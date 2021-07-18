package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"satu", "dua", "tiga"}
	var hasil = filter(data, func(s string) bool {
		return strings.Contains(s, "aa")
	})

	fmt.Printf("hasilnya : %s", hasil)
}

func filter(data []string, callback func(string) bool) []string {

	var result []string
	for _, val := range data {
		if filter := callback(val); filter {
			result = append(result, val)
		}
	}

	return result
}
