package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&n)
	fmt.Println(isPrime(n))
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}