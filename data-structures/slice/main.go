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

	// slice of type int with length of 0 and capacity of 5
	y := make([]int, 0, 8)

	for i := 0; i < 50; i++ {
		y = append(y, i)
		fmt.Println("Length:", len(y), " Capacity:", cap(y), " Value:", y[i])
	}

}