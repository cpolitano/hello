package fundamentals

// have to import packages needed in each package
import "fmt"

// block level scope with inner scope
func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func Closure() {
	fmt.Println("Inside Closure")
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}