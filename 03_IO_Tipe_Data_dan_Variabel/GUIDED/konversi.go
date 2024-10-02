package main

import (
    "fmt"
)

func main() {
    var idr int
    fmt.Print("Masukkan jumlah uang dalam IDR: ")
    fmt.Scan(&idr)

    usd := float64(idr) / 15000.0 
    fmt.Printf("%d IDR setara dengan %.2f USD.\n", idr, usd)
}
