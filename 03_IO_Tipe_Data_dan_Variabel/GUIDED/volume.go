package main

import (
	"fmt"
)

func main() {
	var sisi int
	var hasil int

	fmt.Print("Masukkan panjang sisi kubus (bilangan bulat positif): ")
	fmt.Scanln(&sisi)

	hasil = sisi * sisi * sisi

	fmt.Printf("Volume kubus dengan sisi %d adalah %d\n", sisi, hasil)
}
