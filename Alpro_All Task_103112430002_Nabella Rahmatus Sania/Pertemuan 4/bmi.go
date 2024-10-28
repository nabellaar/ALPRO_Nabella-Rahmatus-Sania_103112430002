package main

import (
	"fmt"
)

func main() {
	var nama string
	var berat, tinggi float64

	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scanln(&berat)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scanln(&tinggi)

	bmi := berat / (tinggi * tinggi) // rumus menghitung BMI

	fmt.Printf("BMI dari %s adalah: %.2f\n", nama, bmi)
}
