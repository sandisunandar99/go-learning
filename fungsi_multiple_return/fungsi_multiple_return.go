package main

import (
	"fmt"
	"math"
)

func main() {
	// ambil fungsi multiple return
	var diameter float64 = 62
	luas, keliling := FungsiMultipeReturun(diameter)

	fmt.Printf("luas =  %f ", luas)
	fmt.Printf("keliling =  %f \n", keliling)

	var NamaAnda = [2]string{"dian ", "sastro"}

	depan, belakang := SiapaNama(NamaAnda)
	fmt.Print(depan)
	fmt.Println(belakang)

}

func FungsiMultipeReturun(d float64) (float64, float64) {

	// hitung luas
	var luas = math.Pi * math.Pow(d/2, 2)

	// keliling
	var keliling = math.Pi * d

	return luas, keliling
}

func SiapaNama(nama [2]string) (depan, belakang string) {
	depan = nama[0]
	belakang = nama[1]
	return
}
