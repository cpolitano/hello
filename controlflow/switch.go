package controlflow

import "fmt"

func Switch() {
	fmt.Println("Switch")
	
	switch "Clare" {
		case "David":
			fmt.Println("Hello David")
		case "Anisa":
			fmt.Println("Hello Anisa")
		case "Clare":
			fmt.Println("Hello Clare")
			fallthrough // will run next block of code as well
		case "Tiffany":
			fmt.Println("Hello Tiffany")
	}
}