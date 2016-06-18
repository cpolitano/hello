package main

import "fmt"

func main() {
	// defer will stop function from executing until the very end of executing the parent function
	// often used to close a file after it is opened and operated on
	defer visit([]int{1, 2, 3, 5}, func(n int) {
		fmt.Println(n)
	})

	xs := filter([]int{1, 1, 2, 3, 5, 8}, func(n int) bool {
		return n > 1
	})

	fmt.Println(xs)

	fmt.Println(factorial(5))
}

//passing a func as an argument
func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
} 

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int // declare xs as slice of int at zero/empty value
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

// common in functional programming
// but that's not really what Go is about
// Go is about clarity and readability, not complexity


// recursion example
// big performance cost, usually better to use loops
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}