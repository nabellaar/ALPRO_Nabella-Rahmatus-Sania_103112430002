package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari float64

	// Meminta input dari pengguna
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&jariJari)

	// Menghitung luas lingkaran
	luasLingkaran := math.Pi * jariJari * jariJari

	// Menampilkan hasil
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", jariJari, luasLingkaran)
}
