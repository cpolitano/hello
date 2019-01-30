package main

import "fmt"

func main() {

	// anonymous functons
	// assigned to a variable => function expression
	// only way to define a function inside another function
	greeting := func() {
		fmt.Println("Aloha")
	}

	greeting()

	greet := makeGreeter()
	fmt.Println(greet())
}

// function that returns a function that returns a string
func makeGreeter() func() string { // type of return is "func() string"
	return func() string { // function that returns a string
		return "Salaam"
	}
}
