package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("Bukan segitiga")
	} else if a == b && b == c {
		fmt.Println("Segitiga sama sisi")
	} else if a == b || b == c || c == a {
		fmt.Println("Segitiga sama kaki")
	} else {
		fmt.Println("Segitiga sembarang")
	}
}
