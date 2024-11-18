// a.) Output program untuk input 80.1
// - Terjadi Kesalahan dalam variabel, variabel nam yang seharusnya adalah variabel nmk (Output program adalah Nilai Mata Kuliah :A)

// b.) Kesalahan Program
// - Kesalahan dalam variabel, untuk output seharusnya menggunakan variabel nmk, bukan nam
// - Program memeriksa kondisi dengan urutan yang salah
// - Program menggunakan if tanpa else, yang membuat program memeriksa kondisi yang sudah terpenuhi.

// c.) Kode program yang sudah diperbaiki :

package main

import "fmt"

func main() {
    var nam float64
    var nmk string
    
    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scan(&nam)

    if nam > 80 {
        nmk = "A"
    } else if nam > 72.5 {
        nmk = "AB"
    } else if nam > 65 {
        nmk = "B"
    } else if nam > 57.5 {
        nmk = "BC"
    } else if nam > 50 {
        nmk = "C"
    } else if nam > 40 {
        nmk = "D"
    } else {
        nmk = "E"
    }

    fmt.Println("Nilai mata kuliah:", nmk)
}
