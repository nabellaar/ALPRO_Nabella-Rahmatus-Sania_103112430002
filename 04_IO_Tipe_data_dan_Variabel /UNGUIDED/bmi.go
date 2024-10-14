package main

import (
	"fmt"
)

func main() {
	var bmi, tinggi float64

	fmt.Print("Masukkan nilai BMI: ") // input nilai BMI
	fmt.Scan(&bmi)

	fmt.Print("Masukkan tinggi badan dalam meter: ") // input tinggi badan
	fmt.Scan(&tinggi)

	berat := bmi * (tinggi * tinggi) // rumus menghitung bb

	fmt.Printf("Berat badan Anda adalah: %.2f kilogram\n", berat) // hasil
}
