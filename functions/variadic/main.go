package main

import "fmt"

func main() {
	total := sum(1, 2, 3, 4)
	fmt.Println(total)

	n := average(22, 46, 72, 83, 56, 29, 96)
	fmt.Println(n)

	data := []float64{25, 57, 34, 29, 79, 83} // composite literal, a slice of floats
	m := sliceAverage(data...)                // Unfurling the Slice: opens slice and passes elements to function as individual args
	fmt.Println(m)
}

// variadic functions with spread operator can take zero or more arguments
// "Lexical element" is spread operator
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // type of x is slice of int

	total := 0 // short variable declaration is preferred
	for index, value := range x {
		fmt.Println("adding value ", value, " at position ", index)
		total += value
	}

	return total
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf) // type is "slice of float64", []float64, which is a different data type from just float64
	total := 0.0            // could also declare as var total float64, which will initialize at zero value
	for _, v := range sf {  // range loops over an iterable. normally i would be where blank is
		total += v
	}
	return total / float64(len(sf))
}

func sliceAverage(sf ...float64) float64 { // could also change to accept a []float64
	total := 0.0
	for _, v := range sf { // _ is index, v is value at index
		total += v
	}
	return total / float64(len(sf))
}
