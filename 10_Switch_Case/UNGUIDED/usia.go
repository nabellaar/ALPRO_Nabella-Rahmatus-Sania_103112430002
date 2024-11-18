package main

import (
	"fmt"
)

func main() {
	var usia int

	fmt.Print("Masukkan usia: ") // input usia oleh pengguna
	fmt.Scan(&usia)

	switch { // menentukan kategori usia menggunakan Switch Case
	case usia >= 0 && usia <= 12:
		fmt.Println("Kategori: Anak-anak")
	case usia >= 13 && usia <= 17:
		fmt.Println("Kategori: Remaja")
	case usia >= 18 && usia <= 64:
		fmt.Println("Kategori: Dewasa")
	case usia >= 65:
		fmt.Println("Kategori: Lansia")
	default:
		fmt.Println("Usia yang dimasukkan tidak valid")
	}
}
