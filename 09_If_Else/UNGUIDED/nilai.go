package main

import (
	"fmt"
)

func main() {
	var nilai float64

	fmt.Print("Masukkan nilai mahasiswa: ") // Input nilai mahasiswa
	fmt.Scanln(&nilai)

	if nilai > 90 {
		fmt.Println("Indeks: A")
	} else if nilai >= 80 && nilai <= 90 {
		fmt.Println("Indeks: AB")
	} else if nilai >= 70 && nilai < 80 {
		fmt.Println("Indeks: B")
	} else if nilai < 70 {
		fmt.Println("Indeks: C")
	} else {
		fmt.Println("Nilai tidak valid")
	}
}
