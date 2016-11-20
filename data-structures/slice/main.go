package main

import "fmt"

func main() {
	// slices are a data structure to store a list of values of a certain type
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
	fmt.Println(greeting[:2]) // prints els up to but not including index 2
	fmt.Println(greeting[3:]) // prints els at index 3 and after
	fmt.Println(greeting[:]) // prints all els in slice
	// the only way to DELETE from a slice
	greeting = append(greeting[:3], greeting[4:]...) // removes "bonjour" at index 3
	fmt.Println(greeting)

	// Multidimensional slices
	// creates a 2D slice, a slice of slice of string
	// can also make a slice of slice of int
	records := make([][]string, 0)
	student1 := make([]string, 4)
	student1[0] = "Aloha"
	student1[1] = "Jones"
	student1[2] = "100.00"
	student1[3] = "93.00"
	records = append(records, student1)
	student2 := make([]string, 4)
	student2[0] = "Tulsi"
	student2[1] = "Smith"
	student2[2] = "86.00"
	student2[3] = "98.00"
	records = append(records, student2)
	fmt.Println(records)

	// Ways to Declare a Slice
	// A slice is a reference type with an underlying data structure
	// A slice has a pointer, length, capacity
	// Shorthand
	// gets underlying data structure, but cannot start referencing index
	// must use append to add elements because slice has no length declared
	studentA := []string{}
	// Var 
	// nothing is made, reference points to nothing
	// initializes as zero value of slice, which is nil
	// var ALWAYS initializes at zero value of the given type
	var studentB []string 
	// Make - the idiomatic way to declare a slice
	// gives the slice length, so it can be added to immediately
	studentC := make([]string, 8)
	fmt.Println(studentA, studentB, studentC)

	// Incrementing a slice
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7 
	fmt.Println(mySlice[0])
	mySlice[0]++
	fmt.Println(mySlice[0])

}