package main

import "fmt"

func main() {
	// slice
	// var x []string 

	var y [58]string // array of type string, which specifies number of elements on init
	// initialized with zero value, it will have 25 empty strings
	// arrays are not used often in Go
	// length is part of type
	// arrays are NOT dynamic -- do not change in size

	for i := 65; i <= 122; i++ {
		y[i-65] = string(i)
	}
	// 65 to 122 represents ascii alphabet

	fmt.Println(y)
	fmt.Println(y[45])

	var z [256]int 
	fmt.Println(len(z))

	for i := 0; i < 256; i++ {
		z[i] = i
	}

	for i, v := range z {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 100 {
			break
		}
	}

}