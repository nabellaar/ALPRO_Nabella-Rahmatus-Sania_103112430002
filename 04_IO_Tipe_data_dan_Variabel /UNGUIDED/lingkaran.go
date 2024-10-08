package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	const phi = math.Pi

	fmt.Print("Masukkan jari-jari lingkaran: ") // Input jari jari lingkaran
	fmt.Scanln(&radius)

	luas := phi * math.Pow(radius, 2) // Menghitung luas lingkaran

	keliling := 2 * phi * radius // Menghitung keliling lingkaran

	fmt.Printf("Luas lingkaran: %.2f\n", luas) // Menampilkan hasil luas lingkaran
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling) // Menampilkan hasil keliling lingkaran
}
