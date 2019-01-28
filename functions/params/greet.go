package params

import "fmt"

// Functions syntax
// func (r receiver) indentifier(parameters) return(s) { code }

func Greet(fname, lname string) { // two parameters, type declared once
	fmt.Println("Hello", fname, lname)
}

// one parameter, one return
func GreetAloha(name string) string {
	return fmt.Sprint("Aloha, ", name)
}

// two parameters, two returns
func GreetMultiple(friend1, friend2 string) (string, bool) {
	a := fmt.Sprint(friend1, " says hello to ", friend2)
	b := true

	return a, b
}
