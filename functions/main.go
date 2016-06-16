package main

import "fmt"
import "github.com/cpolitano/hello/functions/params"
import "github.com/cpolitano/hello/functions/returns"

func main() {
	params.Greet("Avocado", "Smith")
	fmt.Println(returns.Greet("Kale ", "Hutchings"))
	fmt.Println(returns.GreetMultiple("Kale ", "Hutchings"))
}
