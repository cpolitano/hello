package fundamentals

// have to import packages needed in each package
import "fmt"

// package level scope
var x int = 88;

// accessible outside package because capital letter
func Scope() {
	// block level scope
	fmt.Println("Inside Scope")
	y := 17
	fmt.Println(y)

	fmt.Println(x)
	foo()
}

// accessible only in package
func foo() {
	fmt.Println(x)
}