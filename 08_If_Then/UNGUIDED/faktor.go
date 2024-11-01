package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Println("Masukkan dua bilangan bulat positif :")
	fmt.Scan(&x, &y)

	var isXFactorOfY bool
	var isYFactorOfX bool

	// x adalah faktor dari y
	if y%x == 0 {
		isXFactorOfY = true // jika x adalah faktor dari y
	} else {
		isXFactorOfY = false // jika x bukan faktor dari y
	}

	// y adalah faktor dari x
	if x%y == 0 {
		isYFactorOfX = true // jika y adalah faktor dari x
	} else {
		isYFactorOfX = false // jika y bukan faktor dari x
	}

	fmt.Println(isXFactorOfY)
	fmt.Println(isYFactorOfX)
}
