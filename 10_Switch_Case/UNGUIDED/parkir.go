package main

import "fmt"

func main() {
	var jenisKendaraan string
	var durasi float64
	var tarifPerJam int

	fmt.Println("Masukkan jenis kendaraan (motor/mobil/truk):")
	fmt.Scanln(&jenisKendaraan)

	fmt.Println("Masukkan durasi parkir dalam jam:")
	fmt.Scanln(&durasi)

	if durasi < 1 {
		durasi = 1
	}

	switch jenisKendaraan {
	case "motor":
		tarifPerJam = 2000
	case "mobil":
		tarifPerJam = 5000
	case "truk":
		tarifPerJam = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid!")
		return
	}

	totalBiaya := tarifPerJam * int(durasi)

	fmt.Printf("Jenis kendaraan: %s\n", jenisKendaraan)
	fmt.Printf("Durasi parkir: %.0f jam\n", durasi)
	fmt.Printf("Total biaya parkir: Rp %d\n", totalBiaya)
}
