package main

import "fmt"

func main ()  {
	var kata string

	for {
		fmt.Println("Masukkan kata apa saja : ")
		fmt.Scan(&kata)

		if kata == "Telkom" || kata == "TELKOM" || kata == "telkom" {
			fmt.Println("Program selesai")

			break
		} else {
			fmt.Println("Anda mengetik : ", kata)
		}
	}
}