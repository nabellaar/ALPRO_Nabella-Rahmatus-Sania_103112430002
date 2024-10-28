package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Silakan masukkan bilangan bulat positif.")	// output jika bilangan tersebut bukan bilangan bulat positif
		return
	}

	faktorial := 1
	for i := 1; i <= n; i++ {	// menghitung faktorial 
		faktorial *= i
	}

	fmt.Printf("Faktorial dari %d adalah %d\n", n, faktorial)
}
