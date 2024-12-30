package main
import "fmt"

func main ()  {
	var val int

	fmt.Print("Masukkan nilai : ")
	fmt.Scan(&val)

	if val  >= 0 {
		fmt.Println("Positif")
		if val % 2 == 0 {
			fmt.Println("Genap")
		} else {
			fmt.Println("Ganjil")
		}
	} else {
		fmt.Println("Angka adalah negatif")
	}

}