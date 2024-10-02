package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64 // Menyimpan dua variabel yang dimasukkan pengguna
	var operator string // Menyimpan operator yang dipilih oleh pengguna

	fmt.Println("Selamat datang di Kalkulator Sederhana!")
	fmt.Println("Masukkan dua angka: ")

	fmt.Print("Angka pertama: ")
	fmt.Scan(&num1) // Mengambil input angka pertama dari pengguna

	fmt.Print("Angka kedua: ")
	fmt.Scan(&num2) // Mengambil input angka kedua dari pengguna

	fmt.Print("Masukkan operator (+, -, *, /): ")
	fmt.Scan(&operator) // Mengambil input operator aritmatika dari pengguna

	var result float64

	switch operator {
	case "+":
		result = num1 + num2
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, result) // Menampilkan hasil penjumlahan
	case "-":
		result = num1 - num2
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, result) // Menampilkan hasil pengurangan 
	case "*":
		result = num1 * num2
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, result) // Menampilkan hasil perkalian
	case "/":
		if num2 != 0 {
			result = num1 / num2
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, result) // Menampilkan hasil pembagian
		} else {
			fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.") // Menampilkan pesan kesalahan jika pembagian dengan 0
		}
	default:
		fmt.Println("Error: Operator tidak valid.")
	}
}
