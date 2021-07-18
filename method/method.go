package main

import "fmt"

//Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya).

type mahasiswa struct {
	Nama    string
	Jurusan string
}

func (s mahasiswa) Hallo() {
	fmt.Println("Nama :", s.Nama)
	fmt.Println("Jurusan :", s.Jurusan)
}

func (s mahasiswa) Holla(ss string) string {
	return ss
}

func main() {
	var s = mahasiswa{"sandi", "tekom"}
	s.Hallo()

	fmt.Println("Nama :", s.Holla("Dian"))
	fmt.Println("Jurusan :", s.Holla("MTK"))

}
