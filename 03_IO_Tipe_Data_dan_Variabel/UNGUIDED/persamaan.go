package main

import (
	"fmt"
	"math"
)

func calculateX(fx float64) float64 {
	// Persamaan yang diubah: fx = 2/(x+5) + 5
	// Setelah disederhanakan: x = (2 / (fx - 5)) - 5
	if fx <= 5 {
		fmt.Println("Nilai f(x) harus lebih besar dari 5 untuk hasil yang valid.")
		return math.NaN() // Return NaN jika f(x) tidak valid
	}
	
	x := (2 / (fx - 5)) - 5
	return x
}

func main() {
	var fx float64
	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scanln(&fx)

	x := calculateX(fx)

	if !math.IsNaN(x) {
		fmt.Printf("Nilai x adalah: %.4f\n", x)
	}
}
