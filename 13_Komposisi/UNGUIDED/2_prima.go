package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&n)

    if n <= 1 {
        fmt.Println("Bukan prima")
        return
    }

    for i := 2; i < n; i++ {
        if n%i == 0 {
            fmt.Println("Bukan prima")
            return
        }
    }

    fmt.Println("Prima")
}
