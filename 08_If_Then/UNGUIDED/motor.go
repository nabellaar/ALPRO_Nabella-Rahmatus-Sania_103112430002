package main

import (
	"fmt"
)

func main() {
	var jumlahOrang int

	fmt.Print("Masukkan jumlah orang : ") // input jumlah orang
	fmt.Scanln(&jumlahOrang)

	jumlahMotor := jumlahOrang / 2 // menghitung jumlah motor jika orang genap

	if jumlahOrang%2 != 0 {
		jumlahMotor += 1 // menambah 1 motor jika orang ganjil
	}

	fmt.Println("Jumlah motor yang diperlukan: ", jumlahMotor)
}
