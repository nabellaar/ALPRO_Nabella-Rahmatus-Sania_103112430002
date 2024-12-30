package main

import "fmt"

func main() {
	var total int
	var buatKartu bool

	fmt.Scan(&total, &buatKartu)

	if total >= 100000 {
		total -= (total / 10)
	}
	
	if total >= 200000 && buatKartu {
		total -= 75000
	}

	fmt.Println("Kartu?", buatKartu)
	fmt.Println("Diskon?", total >= 100000)
	fmt.Println("Cashback?", buatKartu && total+75000 >= 200000)
	fmt.Println("Rp.", total)
}