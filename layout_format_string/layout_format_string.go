package main

import "fmt"

type student struct {
	name        string
	age         int32
	height      float64
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "sandi",
	age:         17,
	height:      67.6,
	isGraduated: true,
	hobbies:     []string{"touring", "sepedahan"},
}

func main() {
	fmt.Printf("%.2f", data.height)
}
