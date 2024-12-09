package main

import (
	"fmt"
)

// Menghitung jumlah faktor bilangan
func jumlahFaktor(bilangan int) int {
	jumlah := 0
	for i := 1; i < bilangan; i++ { 
		if bilangan%i == 0 { 
			jumlah += i
		}
	}
	return jumlah
}

func main() {
	var bilangan int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&bilangan) 

	if jumlahFaktor(bilangan) == bilangan {
		fmt.Println("Ya, bilangan tersebut adalah bilangan sempurna.")
	} else {
		fmt.Println("Tidak, bilangan tersebut bukan bilangan sempurna.")
	}
}
