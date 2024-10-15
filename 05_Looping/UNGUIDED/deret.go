package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif n: ") // input pengguna
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Harap masukkan bilangan bulat positif!") // memastikan bilangan positif atau bukan
		return
	}

	sum := 0
	for i := 1; i <= n; i++ { // menghitung jumlah deret bilangan
		sum += i
	}

	fmt.Printf("Jumlah deret dari 1 hingga %d adalah: %d\n", n, sum)
}
