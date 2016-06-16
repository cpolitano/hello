package returns

import "fmt"

// two arguments, type declared once
// RETURN is defined as string
func Greet(fname, lname string) string { 
	// Sprint is string print
	return fmt.Sprint(fname, lname)
}

func GreetMultiple(fname, lname string) (string, string) { 
	// Sprint is string print
	return fmt.Sprint(fname, lname), fmt.Sprint(fname)
}
