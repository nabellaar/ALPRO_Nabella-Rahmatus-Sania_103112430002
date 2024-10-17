package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("Masukkan bilangan pertama: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan bilangan kedua: ")
	fmt.Scan(&b)

	hasil := 1

	for i := 0; i < b; i++ {	// menghitung pemangkatan dengan perulangan
		hasil *= a
	}

	fmt.Println("Hasil dari", a, "pangkat", b, "=", hasil)
}
