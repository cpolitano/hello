package main

import "fmt"
import "github.com/cpolitano/hello/functions/params"
//import "github.com/cpolitano/hello/functions/returns"

func main() {
	params.Greet("Avocado", "Smith")
	greet := params.GreetAloha("Dolphins")
	fmt.Println(greet)
	x, y := params.GreetMultiple("Kale", "Coriander") // multiple assignment
	fmt.Println(x, y)
	// fmt.Println(returns.Greet("Kale ", "Hutchings"))
	// fmt.Println(returns.GreetMultiple("Kale ", "Hutchings"))
}
