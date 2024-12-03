package main

import "fmt"

func main ()  {
	var angka int
	angkaRahasia := 5

	for {
		fmt.Println("Masukkan angka dari 1-10 : ")
		fmt.Scan(&angka)

		if angka == angkaRahasia {
			fmt.Println("Selamat tebakan anda benar!")

			break
		} else {
			fmt.Println("Tebakan anda salah, silahkan coba lagi")
		}
	}
}