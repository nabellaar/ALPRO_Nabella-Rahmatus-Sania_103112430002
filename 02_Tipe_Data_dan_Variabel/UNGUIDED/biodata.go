package main

import (
	"fmt"
)

func main() {
	var nama, nim, kelas string

	// Meminta input dari pengguna
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan NIM Anda: ")
	fmt.Scanln(&nim)
	fmt.Print("Masukkan kelas Anda: ")
	fmt.Scanln(&kelas)

	// Menyusun dan menampilkan resume biodata
	resume := fmt.Sprintf("Perkenalkan, saya adalah %s, salah satu mahasiswa Prodi S1-IF dari kelas %s dengan NIM %s.", nama, kelas, nim)
	fmt.Println(resume)
}
