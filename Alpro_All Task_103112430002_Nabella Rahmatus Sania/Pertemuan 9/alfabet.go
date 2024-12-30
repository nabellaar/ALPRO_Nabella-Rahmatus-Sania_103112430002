package main

import (
	"fmt"
)

func main() {
	var kar rune
	
	fmt.Println("Masukkan satu karakter: ")
	fmt.Scanf("%c", &kar)

	if ('a' <= kar && kar <= 'z') || ('A' <= kar && kar <= 'Z') {
		fmt.Println("Alphabet")
	} else {
		fmt.Println("Bukan Alphabet")
	}
}