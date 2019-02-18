package main

import (
	"fmt"
)

func getAverage(m float64, n float64) float64 {
	return (m + n) / 2
}

func getSum(xi ...int) int {
	sum := 0
	for _, v := range xi { // i index not needed, replace with _
		sum += v
	}
	return sum
}

func main() {
	x := getAverage(4.7283, 8.22)
	y := getAverage(8822.3, 2636.2)
	fmt.Println(x)
	fmt.Println(y)

	s := getSum(3, 6, 2, 8)
	t := getSum(44,2,5,1,22)
	fmt.Println(s)
	fmt.Println(t)
}
