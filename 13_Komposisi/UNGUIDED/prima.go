package main

import "fmt"

// Memeriksa apakah bilangan prima
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var limit int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&limit)

	fmt.Println("Bilangan prima dari 1 hingga", limit, "adalah:")
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
