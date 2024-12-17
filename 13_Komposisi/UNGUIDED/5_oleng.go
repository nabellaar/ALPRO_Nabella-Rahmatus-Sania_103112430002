package main

import (
	"fmt"
	"math"
)

func main() {
	var beratKiri, beratKanan float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&beratKiri, &beratKanan)

		if beratKiri < 0 || beratKanan < 0 {
			fmt.Println("Proses selesai.")
			break
		}
		if beratKiri+beratKanan > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		oleng := math.Abs(beratKiri-beratKanan) >= 9
		fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", oleng)
	}
}