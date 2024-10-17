package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat non-negatif: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Bilangan harus non-negatif.") // validasi bilangan positif
		return
	}

	faktorial := 1
	for i := 1; i <= n; i++ {	// menghitung faktorial
		faktorial *= i
	}

	fmt.Println("Faktorial dari", n, "adalah:", faktorial)
}
