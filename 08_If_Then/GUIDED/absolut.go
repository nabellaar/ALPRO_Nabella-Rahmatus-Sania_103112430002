package main 

import "fmt" 

func main() {
	var bilangan int 

	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scan(&bilangan) 

	if bilangan < 0 {
	bilangan = -bilangan}

	fmt.Println(bilangan) }