package main

import (
	"fmt"
)

func main() {
	var saldo, transaksi int

	for {
		fmt.Scan(&transaksi)
		if transaksi == 0 {
			break
		}
		saldo += transaksi
	}

	fmt.Println(saldo)
}
