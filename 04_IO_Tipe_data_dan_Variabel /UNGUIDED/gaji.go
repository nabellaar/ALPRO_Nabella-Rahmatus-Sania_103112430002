package main

import (
	"fmt"
)

func main() {
	var jamKerjaPerMinggu float64
	var upahPerJam float64

	// Meminta input jam kerja dalam seminggu
	fmt.Print("Masukkan jumlah jam kerja per minggu: ")
	fmt.Scanln(&jamKerjaPerMinggu)

	// Meminta input upah per jam
	fmt.Print("Masukkan upah per jam: ")
	fmt.Scanln(&upahPerJam)

	var totalGajiBulanan float64
	var jamNormal float64
	var lembur float64

	// Menghitung jam normal dan jam lembur
	if jamKerjaPerMinggu > 40 {
		jamNormal = 40
		lembur = jamKerjaPerMinggu - 40
	} else {
		jamNormal = jamKerjaPerMinggu
		lembur = 0
	}

	// Menghitung total gaji bulanan (4 minggu)
	totalGajiBulanan = (jamNormal * upahPerJam * 4) + (lembur * 1.5 * upahPerJam * 4)

	// Menampilkan total gaji bulanan
	fmt.Printf("Total gaji bulanan: %.2f\n", totalGajiBulanan)
}
