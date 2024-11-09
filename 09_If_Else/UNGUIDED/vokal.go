package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Print("Masukkan satu huruf: ") // input pengguna
	fmt.Scanln(&input)

	// memeriksa apakah input merupakan huruf vokal 
	if input == "A" || input == "I" || input == "U" || input == "E" || input == "O" || 
		input == "a" || input == "i" || input == "u" || input == "e" || input == "o" {
		fmt.Println("Huruf Vokal")
	} else {
		fmt.Println("Huruf Konsonan") // jika bukan huruf vokal
	}
}
