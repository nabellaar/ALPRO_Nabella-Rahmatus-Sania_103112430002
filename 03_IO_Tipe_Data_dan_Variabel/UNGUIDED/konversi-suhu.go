package main

import (
	"fmt"
)

func main() {
	var fahrenheit float64

	fmt.Print("Masukkan suhu dalam derajat Fahrenheit: ") // Input suhu dalam fahrenheit
	fmt.Scan(&fahrenheit)

	kelvin := (fahrenheit-32) * 5/9 + 273.15 // Konversi Fahrenheit ke Kelvin

	fmt.Printf("Suhu dalam derajat Kelvin: %.2f K\n", kelvin) // Hasil dari Kelvin
}
