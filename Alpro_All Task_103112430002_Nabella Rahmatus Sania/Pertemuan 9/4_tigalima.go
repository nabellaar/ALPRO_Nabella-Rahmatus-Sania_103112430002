package main

import "fmt"

func main() {
    var number int
    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&number)

    if number <= 0 {
        fmt.Println("Masukkan bilangan bulat positif.")
    } else if number%3 == 0 && number%5 == 0 {
        fmt.Println("Kelipatan 3 dan 5")
    } else if number%3 == 0 {
        fmt.Println("Kelipatan 3")
    } else if number%5 == 0 {
        fmt.Println("Kelipatan 5")
    }
}
