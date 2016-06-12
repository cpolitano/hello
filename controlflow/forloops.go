package controlflow

import "fmt"

var a int = 10

// Go equivalent of while loop
func while() {
	i := 0
	for i < a {
		fmt.Println(i)
		i++
	}
}

func forbreak() {
	i := 0
	for {
		i++
		if i % 2 == 0 {
			continue 
			// end this iteration of loop
			// code below is not executed
		}
		fmt.Println(i)
		if i > 20 {
			break // break out of loop
		}
	}
}

func forloop() {
	for i := 50; i < 140; i++ { // init, condition, post
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}

func runes() {
	foo := 'a' // denotes RUNE, not a string
	fmt.Println(foo) // will print 97, the utf-8 byte representing lowercase a
	fmt.Printf("%T \n", foo) // %T prints data type of foo, INT32
	fmt.Printf("%T \n", rune(foo))
}

func Loops() {
	fmt.Println("Loops")
	runes()
}