package main

import (
    "fmt"
)

func main() {
    var nilai int

    fmt.Print("Nilai ujian akhir: ") // input nilai ujian pengguna
    fmt.Scan(&nilai)

    if nilai >= 70 {
        fmt.Println("Lulus") // jika nilai lebih dari sama dengan 70
    } else {
        fmt.Println("Tidak Lulus") // jika nilai lebih kecil dari 70
    }
}
