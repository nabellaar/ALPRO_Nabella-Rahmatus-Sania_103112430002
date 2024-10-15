package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // seed untuk pengecekan angka

	target := rand.Intn(100) + 1	// menghasilkan angka acak antara 1 hingga 100
	var guess int
	maxAttempts := 5

	fmt.Println("Selamat datang di permainan tebak angka!")
	fmt.Println("Saya telah memilih angka antara 1 hingga 100.")
	fmt.Println("Kamu punya 5 kesempatan untuk menebak angka yang benar.")

	for attempts := 1; attempts <= maxAttempts; attempts++ {	// loop hingga pengguna berhasil menebak atau kesempatan habis
		fmt.Printf("Tebakan #%d: Masukkan angka: ", attempts)
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Terlalu kecil!")
		} else if guess > target {
			fmt.Println("Terlalu besar!")
		} else {
			fmt.Printf("Selamat! Kamu berhasil menebak angka %d dalam %d percobaan.\n", target, attempts)
			return
		}
	}

	fmt.Printf("Maaf, kamu kehabisan kesempatan. Angka yang benar adalah %d.\n", target) // pesan jika salah menebak hingga akhir
}
