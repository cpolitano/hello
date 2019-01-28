package fundamentals

import "fmt"

const metersToYards float64 = 1.09361

func convertMetersToYards(z *float64) { // takes pointer to float as parameter
	*z = *z * metersToYards // value at pointer z times meters to yards conversion
}

/* Why Pointers
very performant because instead of passing values around like in JS
you can just pass the pointers, and still have access to the values
everything in Go is "Pass by value" - NOT pass by reference, pass by copy (JS)

Example 1 - you have a large piece of data.
instead of passing around all that data into different functions,
you can access it by pointer, the address at which it is stored

Example 2 - you want the address to stay the same for a given value,
so instead just change the value via pointer
*/

func Memory() {

	a := 44
	fmt.Println("a - ", a)
	fmt.Printf("type of a - %T\n", a) // print type of a (type int)

	// &something "ampersand operator" indicates address in memory where value is stored
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("type of &a - %T\n", &a) // print type of &a (*int, type pointer to int)

	var b *int = &a // b is type Pointer to Int, points to memory address of a
	fmt.Println("b - pointer to a's memory address", b)

	// *int - star is part of type pointer to int
	// *b - star operator indicates deferencing address, which shows value at that address
	fmt.Println("value at *b -", *b) // dereferences b, returns value at that memory address

	*b = 50 // reassign value at that memory address
	fmt.Println("new value of a -", a)

	var distance float64
	fmt.Print("Enter meters run: ")
	fmt.Scan(&distance)
	// scans for text input from stdin and saves it to
	// memory address of 'distance'
	fmt.Println(distance, "meters")
	convertMetersToYards(&distance)
	fmt.Println(distance, "yards")
}
