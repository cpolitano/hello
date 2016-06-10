package fundamentals

import "fmt"

const metersToYards float64 = 1.09361

func Memory() {
	
	a := 44

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a, "\n")

	var meters float64
	fmt.Print("Enter meters run: ")
	fmt.Scan(&meters) 
	// scans for text input from stdin and saves it to 
	// meters' memory address
	yards := meters * metersToYards
	fmt.Println(meters, "meters is", yards, "yards.\n")
}