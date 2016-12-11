package main

import (
	"fmt";
	"sort";
)

func main() {
	// sort in order
	type people []string
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)

	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(s)
	fmt.Println(s)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Ints(n)
	fmt.Println(n)

	// sort in reverse order 
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println(studyGroup)

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}
