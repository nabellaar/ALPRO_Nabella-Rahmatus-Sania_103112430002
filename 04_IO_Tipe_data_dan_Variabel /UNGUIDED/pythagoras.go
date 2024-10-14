package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by, cx, cy float64

	fmt.Println("Masukkan koordinat titik A (x y):") // input titik A
	fmt.Scan(&ax, &ay)
	fmt.Println("Masukkan koordinat titik B (x y):") // input titik B
	fmt.Scan(&bx, &by)
	fmt.Println("Masukkan koordinat titik C (x y):") // input titik C
	fmt.Scan(&cx, &cy)
    
	var sisiAB, sisiBC, sisiCA, terpanjang float64
	sisiAB = math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2)) // menghitung sisi AB dengan teorema Phythagoras
	sisiBC = math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2)) // menghitung sisi BC dengan teorema Phythagoras
	sisiCA = math.Sqrt(math.Pow(ax-cx, 2) + math.Pow(ay-cy, 2)) // menghitung sisi CA dengan teorema Phythagoras

	terpanjang = sisiAB			
	if sisiBC > terpanjang {	// mengecek sisi BC lebih panjang dari AB atau tidak
		terpanjang = sisiBC		// output jika sisi BC lebih panjang
	}
	if sisiCA > terpanjang {	// apakah sisi CA lebih panjang atau tidak
		terpanjang = sisiCA		// output jika sisi CA lebih panjang
	}

	fmt.Printf("Panjang sisi terpanjang pada segitiga adalah: %.2f\n", terpanjang)
}