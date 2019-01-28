package fundamentals

import "fmt"

const k string = "Hello World"
const j int = 50

// can also declare vars of same type on the same line
var p, q int

func Variables() {
	fmt.Println("Variables")
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", k)
	fmt.Printf("%v \n", j)

	p = 40
	q = 41

	// multiple assignments
	p, q = 42, 43
	fmt.Println(p, q)
}
