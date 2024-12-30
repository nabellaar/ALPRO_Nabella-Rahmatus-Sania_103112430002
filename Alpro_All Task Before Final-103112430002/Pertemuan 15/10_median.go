package main

import "fmt"

func main() {
	var y, total, angka int

	fmt.Print("Masukkan nilai y: ")
	fmt.Scanln(&y)

	for i := 0; i < 9; i++ {
		fmt.Scan(&angka)
		total += angka
	}

	if total >= y*5 {
		fmt.Println("Median bernilai" , y)
	} else {
		fmt.Println("Median bernilai 0")
	}
}
