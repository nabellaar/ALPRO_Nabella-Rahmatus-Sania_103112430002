package main 
import "fmt"

func main(){
var j, bil1, bil2 int
var hasil int 

fmt.Print("Masukkan angka pertama : ")
fmt.Scan(&bil1)
fmt.Print("Masukkan angka kedua : ")
fmt.Scan(&bil2)

hasil = 0
for j = 1; j <= bil2; j++ {
    hasil = hasil + bil1}
fmt.Print("Hasil dari perkalian dua angka tersebut adalah : ")
fmt.Println(hasil) }