package main

import "fmt"

func main() {
	var input int

	fmt.Print("Masukkan jumlah baris segitiga: ") // meminta input dari pengguna
	fmt.Scan(&input)

	for i := 1; i <= input; i++ {	// perulangan untuk mencetak segitiga bintang
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
