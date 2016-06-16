package exercises

import "fmt"

func PrintEvens() {
	
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
}
