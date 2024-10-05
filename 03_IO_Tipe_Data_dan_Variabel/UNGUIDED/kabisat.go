package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	// Tahun kabisat: habis dibagi 400 atau (habis dibagi 4 dan tidak habis dibagi 100)
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	var year int
	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&year)

	if isLeapYear(year) {
		fmt.Println("Tahun", year, "adalah tahun kabisat (true).")
	} else {
		fmt.Println("Tahun", year, "bukan tahun kabisat (false).")
	}
}
