package main
import "fmt"

func main() {
	var pita, bunga string
	var count, n int

	fmt.Print("N: ")
	fmt.Scanln(&n)

	if n > 0 {
		for i := 1; i <= n; i++ {
			fmt.Printf("Bunga %d: ", i)
			fmt.Scanln(&bunga)
			if pita == "" {
				pita = bunga
			} else {
				pita += " – " + bunga
			}
			count++
		}
		pita += " -"
	}
	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", count)
}