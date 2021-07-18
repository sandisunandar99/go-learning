package main

import "fmt"

type mahasiswa struct {
	nama  string
	absen int
}

type kuliah struct {
	mapel string
	nilai int
}

type ujian struct {
	ke int
	kuliah
}

func main() {
	var m mahasiswa

	m.nama = "sandi"
	m.absen = 10

	fmt.Println("Nama :", m.nama)
	fmt.Println("Absen :", m.absen)

	var m1 = kuliah{"mtk", 1}
	var m2 = kuliah{"ktk", 100}

	fmt.Println("Mapel :", m1.mapel)
	fmt.Println("Nilai :", m1.nilai)
	fmt.Println("Mapel :", m2.mapel)
	fmt.Println("Nilai :", m2.nilai)

	// variabale object pointer
	var m3 *kuliah = &m2
	fmt.Println("Mapel :", m3.mapel)
	fmt.Println("Nilai :", m3.nilai)

	// embed struct atau menggabungkan struct kedalam struct yang lain

	var u ujian
	u.ke = 1
	u.mapel = "asd"
	u.nilai = 100

	fmt.Println("hasil ujian :", u)

}
