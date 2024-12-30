package main

import "fmt"

func main() {
	var a, b, c int
	var tipe string

	fmt.Println("Masukkan tipe kendaraan (A, B, C), akhiri dengan karakter tidak valid:")
	for {
		fmt.Scan(&tipe)
		if tipe == "A" {
			a++
		} else if tipe == "B" {
			b++
		} else if tipe == "C" {
			c++
		} else {
			break
		}
	}

	fmt.Println("Tipe A =", a)
	fmt.Println("Tipe B =", b)
	fmt.Println("Tipe C =", c)
}
