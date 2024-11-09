package main

import (
	"fmt"
)

func main() {
	var umur int
	var kewarganegaraan string

	fmt.Print("Umur Anda: ") // input umur pengguna
	fmt.Scan(&umur)

	fmt.Print("Kewarganegaraan Anda (WNI/WNA): ") // input kewarganegaraan
	fmt.Scan(&kewarganegaraan)

	// mengecek apakah user memenuhi syarat untuk mengikuti Pemilu
	if umur >= 17 && (kewarganegaraan == "WNI" || kewarganegaraan == "wni") {
		fmt.Println("Anda bisa mengikuti pemilu")
	} else {
		if umur < 17 {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena umur Anda kurang dari 17 tahun.")
		}
		if kewarganegaraan != "WNI" && kewarganegaraan != "wni" {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena Anda bukan WNI.")
		}
	}
}
