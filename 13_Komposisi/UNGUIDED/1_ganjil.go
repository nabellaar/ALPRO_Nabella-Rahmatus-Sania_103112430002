package main

import "fmt"

func main() {
    var n int

    fmt.Scan(&n)

    count := 0

    for i := 1; i <= n; i++ {
        if i%2 != 0 { 
            count++
        }
    }
    fmt.Println("Banyaknya bilangan ganjil adalah" , count)
}
