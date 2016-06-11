package fundamentals

import "fmt"

const metersToYards float64 = 1.09361

func convertMetersToYards(z *float64) {
	*z = *z * metersToYards // value at pointer z times meters to yards conversion
} 

func Memory() {
	
	a := 44
	var b *int = &a // b is type Pointer to Int, points to memory address of a

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Println("b - pointer to a's memory address", b)
	fmt.Println(*b) // dereferences b, returns value at that memory address
	*b = 50 // reassign value at that memory address
	fmt.Println(a)
	// very performant because instead of passing values around like in JS
	// you can just pass the pointers, and still have access to the values
	// everything in Go is "Pass by value"

	var meters float64
	fmt.Print("Enter meters run: ")
	fmt.Scan(&meters) 
	// scans for text input from stdin and saves it to 
	// meters' memory address
	fmt.Println(meters, "meters")
	convertMetersToYards(&meters)
	fmt.Println(meters, "yards")
}