package main

import "fmt"

func main() {
	var T, V, totalAir int

	fmt.Print("Masukkan kapasitas tanki (T): ")
	fmt.Scan(&T)

	for {
		fmt.Print("Masukkan volume ember (V) atau 0 untuk berhenti: ")
		fmt.Scan(&V)

		if V == 0 {
			break
		}

		totalAir += V
		fmt.Println(totalAir >= T)
	}
}
