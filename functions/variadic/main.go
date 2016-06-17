package main

import "fmt"

func main() {
	n := average(22, 46, 72, 83, 56, 29, 96)
	fmt.Println(n)
	data := []float64{25, 57, 34, 29, 79, 83} //a slice
	m := sliceAverage(data...) // opens slice and passes elements to function as individual args
	fmt.Println(m)
}

// variadic functions with spread operator can take zero or more arguments
func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf) // type is "slice of float64", []float64, which is a different data type from just float64
	total := 0.0 // could also declare as var total float64, which will initialize at zero value
	for _, v := range sf { // range loops over an iterable. normally i would be where blank is
		total += v
	}
	return total / float64(len(sf))
}

func sliceAverage(sf ...float64) float64 { // could also change to accept a []float64
	total := 0.0 
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
