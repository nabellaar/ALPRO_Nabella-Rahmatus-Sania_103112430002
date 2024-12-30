package main
import "fmt"

func main ()  {
	var bilangan int

	fmt.Print("Masukkan sebuah bilangan :")
	fmt.Scan(&bilangan)

	if bilangan < 0 {
		bilangan = -bilangan
	}

	fmt.Println(bilangan)

	// input := 10

	// if input > 10 {
	// 	fmt.Println(input, "lebih besar dari 10")
	// } else if input >= 10 {
	// 	fmt.Println(input, "sama dengan 10")
	// } else {
	// 	fmt.Println(input, "lebih kecil dari 10")
	// }
}
