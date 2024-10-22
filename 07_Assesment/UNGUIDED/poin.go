package main

import (
	"fmt"
)

func main() {
	var jumlahBarang int

	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahBarang)

	totalPoin := 0

	// memeriksa jumlah barang yang dibeli 
	if jumlahBarang <= 5 {
		totalPoin = jumlahBarang * 10	// jika jumlah barang kurang dari atau sama dengan 5
	} else {
		totalPoin = 5*10 + (jumlahBarang-5)*(10+5)	// jika jumlah barang lebih dari 5
	}

	fmt.Println("Poin yang didapatkan:", totalPoin, "poin")
}
