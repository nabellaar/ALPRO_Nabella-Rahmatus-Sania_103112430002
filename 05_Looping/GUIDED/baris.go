package main 
import "fmt"

func main() { 
	var a, b int
	var j int
	fmt.Print("Masukkan angka pertama : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan angka kedua : ")
	fmt.Scan(&b)

	for j = a; j <=b; j++ {
	fmt.Print(j," ") 
	}
}