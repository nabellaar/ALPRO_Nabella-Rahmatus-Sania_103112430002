package main

import "fmt"

func main() {
    var n int

    fmt.Scan(&n) 

    if n%2 == 0 {
        fmt.Println("Input harus bilangan ganjil!")
        return
    }
  
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ { 
            if j == i || j == n-i-1 {
                fmt.Print(i + 1)
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println() 
    }
}
