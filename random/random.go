package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(10)

	fmt.Println("random 1 :", rand.Int())
	fmt.Println("random 2 :", rand.Int())
	fmt.Println("random 3 :", rand.Int())
}
