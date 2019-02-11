package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 13, 8, 3, 75, 34, 7, 2, 52, 25, 26}
	xs := []string{"Mango", "Papaya", "Aguacate", "Chayote", "Zapote"}
	xf := []float64{5.098, 2.019, 4.00003, 1.9, 1.922, 3.245667}

	sort.Ints(xi) // no reassignment needed, changes slice in place
	// a SLICE is a pointer, length, and capacity

	sort.Strings(xs)
	sort.Float64s(xf)

	fmt.Println(xi)
	fmt.Println(xs)
	fmt.Println(xf)
}
