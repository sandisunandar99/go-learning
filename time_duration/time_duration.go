package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	time.Sleep(5 * time.Second)

	duration := time.Since(start)

	fmt.Println("dalam detik :", duration.Seconds())
	fmt.Println("dalam menit :", duration.Minutes())
	fmt.Println("dalam jam :", duration.Hours())
}
