package main

import (
	"fmt"
)

func main() {
	var celsius float64

	// Input suhu dalam Celsius
	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scanln(&celsius)

	// Konversi suhu
	fahrenheit := (celsius * 9 / 5) + 32         // Konversi ke Fahrenheit
	reamur := celsius * 4 / 5                     // Konversi ke Reamur
	kelvin := celsius + 273.15                    // Konversi ke Kelvin

	// Output hasil konversi
	fmt.Printf("Suhu dalam Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Suhu dalam Reamur: %.2f\n", reamur)
	fmt.Printf("Suhu dalam Kelvin: %.2f\n", kelvin)
}
