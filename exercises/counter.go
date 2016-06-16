package exercises

import "fmt"

func Counter() {
	counter := 0
	
	for i := 0; i <= 100; i++ {
		if i % 3 == 0 {
			counter += 1
		} else if i % 5 == 0 {
			counter += 1
		} 
	}

	fmt.Println(counter)
}
