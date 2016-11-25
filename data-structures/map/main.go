package main

import "fmt"

func main() {
	// maps are for key/value storage
	// reference type pointing to underlying hash table / dictionary
	// value of uninitialized map is nil
	m := make(map[string]int)
	m["f0"] = 1
	m["f1"] = 1
	m["f2"] = 2
	m["f3"] = 3
	m["f4"] = 5
	m["f5"] = 8
	fmt.Println("map:", m)
}