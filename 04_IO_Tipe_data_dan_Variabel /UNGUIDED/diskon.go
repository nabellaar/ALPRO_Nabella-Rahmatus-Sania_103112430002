package main
import "fmt"
func main () {
	var hargaBelanja, diskon, totalAkhir int

	fmt.Print("Masukkan total harga belanja : ")
	fmt.Scan(&hargaBelanja)

	fmt.Print("Masukkan diskon dalam persen : ")
	fmt.Scan(&diskon)

	totalAkhir = hargaBelanja - (hargaBelanja * (diskon / 100))
	fmt.Println("Total harga belanja setelah diskon : " , totalAkhir)

}