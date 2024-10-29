package main

import (
    "fmt"
)

func main() {
    var angka int

    fmt.Print("Masukkan sebuah angka: ") // input angka dari pengguna
    fmt.Scan(&angka)

    if angka%2 == 0 {
        fmt.Println("Angka adalah Genap.") // mengecek apakah angka habis dibagi 2 atau tidak
    } else {
        fmt.Println("Angka adalah Ganjil.")
    }
}
