package main

import "fmt"

func main() {
    var num, d, counter int

    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&num)

    counter = 0 

    for num > 0 {
        d = num % 10
        if d%2 == 0 && d != 0 { 
            counter = counter + 1 
        }
        num = num / 10 
    }

    fmt.Println("Jumlah digit genap ", counter)
}
