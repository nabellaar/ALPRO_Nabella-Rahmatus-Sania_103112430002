package main

import (
	"fmt"
)

func main() {
	var sisi float64 = 27
	
	var keliling = 4 * sisi // Rumus Keliling persegi

	var luas = sisi * sisi // Rumus Luas Persegi

	fmt.Printf("Keliling alun-alun: %.2f meter\n", keliling)
	fmt.Printf("Luas alun-alun: %.2f meter persegi\n", luas)
}
