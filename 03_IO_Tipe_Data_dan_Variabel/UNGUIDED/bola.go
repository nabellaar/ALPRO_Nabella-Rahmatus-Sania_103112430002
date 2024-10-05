package main

import (
	"fmt"
	"math"
)

func main() {
	// Konstanta nilai pi
	const pi = 3.1415926535

	// Input dari pengguna
	var r int
	fmt.Print("Jejari = ")
	fmt.Scanln(&r)

	// Menghitung volume bola: volume = 4/3 * pi * r^3
	volume := (4.0 / 3.0) * pi * math.Pow(float64(r), 3)

	// Menghitung luas permukaan bola: luas = 4 * pi * r^2
	luas := 4 * pi * math.Pow(float64(r), 2)

	// Output hasil perhitungan
	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f.\n", r, volume, luas)
}
