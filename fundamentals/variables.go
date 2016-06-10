package fundamentals

import "fmt"

const k string = "Hello World"
const j int = 50

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
	fmt.Printf(k, j, "\n")
}