package main

import (
	"fmt"
	"os/exec"
)

/**
	Exec digunakan untuk eksekusi perintah command line lewat kode program.
**/

func main() {
	var output, _ = exec.Command("go", "version").Output()
	fmt.Printf(string(output))
}
