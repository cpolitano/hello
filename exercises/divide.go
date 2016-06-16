package exercises

import "fmt"

func DivideTwoNumbers() {
	
	var largeNum int
	var smallNum int

	fmt.Print("Enter a large number: ")
	fmt.Scan(&largeNum)
	fmt.Print("Enter a smaller number: ")
	fmt.Scan(&smallNum)
	fmt.Println(largeNum, " / ", smallNum, " = ", largeNum / smallNum)
}