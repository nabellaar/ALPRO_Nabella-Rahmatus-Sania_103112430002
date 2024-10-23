package main 

import "fmt"

func main() {
	var x, y, j, jumlah int 
	
	fmt.Print("Masukkan bilangan pertama : ")
	fmt.Scan(&x)

	fmt.Print("Masukkan bilangan kedua : ")
	fmt.Scan(&y)

	jumlah = 0

	for j = x; j <= y; j += 1 {
			jumlah = jumlah + j}

	fmt.Println("Hasil penjumlahan = ", jumlah) 
}