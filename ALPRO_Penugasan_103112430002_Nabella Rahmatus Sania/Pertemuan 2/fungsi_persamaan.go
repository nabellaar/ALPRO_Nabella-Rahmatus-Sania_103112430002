package main

import (
    "fmt"
)

func hitungF(x, y int) float64 {
    return (1.0 / float64(3*x*x+10)) + 10*float64(y) + 7
}

func main() {
    var x, y int

    fmt.Print("Masukkan nilai x: ")
    fmt.Scan(&x)
    
    fmt.Print("Masukkan nilai y: ")
    fmt.Scan(&y)

    hasil := hitungF(x, y)
    fmt.Printf("Nilai f(%d, %d) = %.6f\n", x, y, hasil)
}