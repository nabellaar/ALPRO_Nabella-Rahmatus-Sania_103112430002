package main

import "fmt"

func main() {
    var token string
    fmt.Print("Masukkan token: ")
    fmt.Scan(&token)

    for token != "12345abcde" {
        fmt.Println("Token salah")
        fmt.Print("Masukkan token: ")
        fmt.Scan(&token)
    }

    fmt.Println("Selamat, Anda berhasil login!")
}
