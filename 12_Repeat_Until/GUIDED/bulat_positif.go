package main

import "fmt"

func main ()  {
	var angka int
	var lanjutLooping bool

	for lanjutLooping = true; lanjutLooping; {
		fmt.Println("Masukkan angka : ")
		fmt.Scan(&angka)

		lanjutLooping = angka <= 0
	}

	fmt.Println(angka, "adalah bilangan bulat positif")

}