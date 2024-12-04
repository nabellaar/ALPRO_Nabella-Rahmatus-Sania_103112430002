package main

import "fmt"

func main ()  {
	var kata string
	var jumlahUlang int

	fmt.Scan(&kata, &jumlahUlang)

	penghitung := 0
	selesai := false

	for !selesai {
		fmt.Println(kata)
		penghitung ++
		selesai = penghitung >= jumlahUlang
	}
}


