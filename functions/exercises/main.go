package main

import "fmt"

func main() {
	fmt.Println(isEven(5))

	getHalf := func(x int) (int, bool) {
		return x / 2, x % 2 == 0
	}

	h, even := getHalf(6)
	fmt.Println(h, even)

	fmt.Println(greatest(3, 5, 8, 21, 6, 14, 2, 9, 4, 53, 31))
	fmt.Println(variadic(1, 2))
	fmt.Println(variadic(1, 2, 3))
	aSlice := []int{1, 2, 4, 6}
	fmt.Println(variadic(aSlice...))
	fmt.Println(variadic())
}

func isEven(x int) (int, bool) {
	return x / 2, x % 2 == 0
}

func greatest(numbers ...int) int {
	var x int
	for _, v := range numbers {
		if v > x {
			x = v
		}
	}
	return x
}

func variadic(numbers ...int) int {
	var total int 
	for _, v := range numbers {
		total += v
	}
	return total
}
