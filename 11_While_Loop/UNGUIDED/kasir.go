package main

import (
	"fmt"
)

func main() {
	var total float64
	var pilihan string

	for {
		fmt.Println("\n=== Kasir Sederhana ===")
		fmt.Println("1. Tambah Barang") // kode untuk tambah barang
		fmt.Println("2. Lihat Total")   // kode untuk melihat total
		fmt.Println("3. Selesai")       // menyelesaikan program
		fmt.Print("Pilih menu (1/2/3): ")
		fmt.Scanln(&pilihan)

		// kondisi untuk tambah barang
		if pilihan == "1" {
			var namaBarang string
			var hargaBarang float64

			fmt.Print("Masukkan nama barang: ")
			fmt.Scanln(&namaBarang)
			fmt.Print("Masukkan harga barang: ")
			fmt.Scanln(&hargaBarang)

			total += hargaBarang
			fmt.Printf("Barang '%s' dengan harga %.2f berhasil ditambahkan.\n", namaBarang, hargaBarang)

			// kondisi untuk melihat total barang
		} else if pilihan == "2" {
			fmt.Printf("Total belanja saat ini: %.2f\n", total)

			// kondisi untuk menyelesaikan program
		} else if pilihan == "3" {
			fmt.Printf("Total akhir: %.2f\n", total)
			fmt.Println("Terima kasih telah menggunakan kasir sederhana!")
			break

			// kondisi jika input tidak valid
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
