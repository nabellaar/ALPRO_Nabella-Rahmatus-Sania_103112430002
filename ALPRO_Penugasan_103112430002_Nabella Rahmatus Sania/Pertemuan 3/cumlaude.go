package main

import "fmt"

func IsCumlaude(semester int, nilaiEPrT int) bool { // mengecek apakah mahasiswa tersebut cumlaude, berdasarkan dengan
    return semester <= 8 && nilaiEPrT >= 500        // jumlah semester dan nilai EPrT
}

func main() {
    var semester, nilaiEPrT int

    fmt.Print("Masukkan jumlah semester: ")
    fmt.Scan(&semester)
    fmt.Print("Masukkan nilai EPrT: ")
    fmt.Scan(&nilaiEPrT)

    if IsCumlaude(semester, nilaiEPrT) {
        fmt.Println("Mahasiswa cumlaude: true") // jika mahasiswa memenuhi syarat maka output bernilai true
    } else {
        fmt.Println("Mahasiswa cumlaude: false") // jika tidak memenuhi syarat maka output bernilai false
    }
}
