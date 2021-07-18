package main

import (
	"fmt"
	"math"
)

/**
interface merukan sebuah tipe data,
di dalam interface terdapat kumpulan definisi method yang tidak memiliki nilai/ isi
*/

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

func main() {
	var hitungLingkaran hitung

	hitungLingkaran = lingkaran{24.2}
	fmt.Println("=== HITUNG LINGKARAN ===")
	fmt.Println("Luas : ", hitungLingkaran.luas())
	fmt.Println("Keliling :", hitungLingkaran.keliling())
	fmt.Println("Jari-jari : ", hitungLingkaran.(lingkaran).jariJari())
}
