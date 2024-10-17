package main 
import "fmt"

func main() {
	var m, alas, tinggi, n int 
	var luas float64 

	fmt.Print("Masukkan jumlah perulangan : ")
	fmt.Scan(&n)

	for m = 1; m <=n; m++ {
	fmt.Print("Masukkan alas segitiga : ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi segitiga : ")
	fmt.Scan(&tinggi)
	luas = 0.5 * float64(alas * tinggi) 
	fmt.Println(luas)
} }