package main 
import "fmt" 
func main () {
	var BeratBadan, TinggiBadan, bmi float64
	fmt.Scan(&BeratBadan, &TinggiBadan)
	
	bmi = BeratBadan / (TinggiBadan * TinggiBadan)

	fmt.Printf("%.2f", bmi)
}