package main
import "fmt"
func main () {
	var bilangan, d1, d2, d3 int

	fmt.Print("Masukkan bilangan tiga digit : ")
	fmt.Scan(&bilangan)

	d1 = bilangan / 100			// untuk mengambil digit pertama
	d2 = (bilangan % 100) / 10	// untuk mengambil digit kedua
	d3 = bilangan % 10			// untuk mengambil digit ketiga

	fmt.Println(d1 <= d2 && d2<= d3)	// menggunakan operator AND untuk mengecek
}