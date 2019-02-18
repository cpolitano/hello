package main

import (
	"fmt"
)

func getAverage(m float64, n float64) float64 {
	return (m + n)/2
}

func main() {
	x := getAverage(4.7283, 8.2209)
	y := getAverage(8822.3, 2636.2)
	fmt.Println(x)
	fmt.Println(y)
}
