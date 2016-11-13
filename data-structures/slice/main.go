package main

import "fmt"

func main() {
	// slices are a reference type
	// represent an underlying array
	x := []int{1, 1, 2, 3, 5, 8} 
	// unlike arrays, slices are dynamic--their length can change
	fmt.Printf("%T\n", x)
	fmt.Println(x)

	// portion of a slice
	// returns elements starting at position 3 and up to but not including 5
	fmt.Println(x[3:5]) 

	// slice of type int with length of 0 and capacity of 8
	// make([]type, length, capacity)
	// make([]type, length & capcity)
	y := make([]int, 0, 8)

	for i := 0; i < 50; i++ {
		y = append(y, i)
		// as more values are added, capacity of underlying array is exceeded
		// Go creates a new array with double the size
		// this is a performance cost
		fmt.Println("Length:", len(y), " Capacity:", cap(y), " Value:", y[i])
	}

	// more ways to slice
	greeting := []string{
		"hello",
		"ciao",
		"bonjour",
		"hola",
		"hi",
		"gutten morgen",
		"aloha",
	}

	fmt.Println(greeting[1:2]) // prints el at index 1
	fmt.Println(greeting[:2]) // prints els up to index 2
	fmt.Println(greeting[3:]) // prints els at index 3 and after
	fmt.Println(greeting[:]) // prints all els in slice

}