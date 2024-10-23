package main

import "fmt"

func main() {
	var N int

	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		fmt.Print(i * i, " ")
	}
	fmt.Println() 
}
