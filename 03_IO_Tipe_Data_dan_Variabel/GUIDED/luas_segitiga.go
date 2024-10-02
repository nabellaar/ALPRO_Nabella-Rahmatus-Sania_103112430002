package main

import (
	"fmt"
)

func main() {
	var alas, tinggi, luas float64

	fmt.Print("Masukkan panjang alas segitiga (bilangan positif): ")
	fmt.Scanln(&alas)
	fmt.Print("Masukkan tinggi segitiga (bilangan positif): ")
	fmt.Scanln(&tinggi)

	luas = 0.5 * alas * tinggi

	fmt.Println("Luas segitiga adalah", luas)
}
