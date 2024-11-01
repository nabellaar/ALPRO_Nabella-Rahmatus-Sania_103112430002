package main

import (
	"fmt"
)

func main() {
	var angka int

	fmt.Print("Masukkan sebuah angka : ") // input angka 
	fmt.Scan(&angka)

	if angka < 0 && angka%2 == 0 { // jika angka termasuk genap negatif
		fmt.Println("bilangan adalah genap negatif") 
	} else { 
		fmt.Println("bilangan adalah bukan genap negatif") // jika angka TIDAK termasuk genap negatif
	}
}
