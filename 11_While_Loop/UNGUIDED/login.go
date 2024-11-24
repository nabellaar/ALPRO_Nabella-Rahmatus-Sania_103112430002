package main

import "fmt"

func main() {
    correctPassword := "12345" // password yang benar

    for attempts := 1; attempts <= 3; attempts++ { // batas mencoba maksimal 3x menggunakan perulangan
        var input string
        fmt.Print("Masukkan password: ") // input pengguna
        fmt.Scanln(&input)

        if input == correctPassword {
            fmt.Println("Login berhasil!") // output jika password benar
            return
        }
        fmt.Println("Password salah.") // output jika password salah
    }
    fmt.Println("Login ditolak.") // output jika sudah mencapai batas percobaan
}
