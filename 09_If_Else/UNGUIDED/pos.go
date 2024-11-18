package main

import (
	"fmt"
)

func main() {
	var beratParsel int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratParsel)

	beratKg := beratParsel / 1000
	sisaGram := beratParsel % 1000

	biayaPerKg := 10000
	biayaTotal := beratKg * biayaPerKg
	biayaSisa := 0

	if beratKg > 10 {
		biayaSisa = 0
	} else {
		if sisaGram >= 500 {
			biayaSisa = sisaGram * 5
		} else {
			biayaSisa = sisaGram * 15
		}
	}

	totalBiaya := biayaTotal + biayaSisa

	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaTotal, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
