package main

import "fmt"

// in Go we create types, not classes
// a struct is a sequence of named elements called fields
// fields have a name and a type

type Person struct {
	first string
	last string
	age int
}

type Singer struct {
	Person
	genre string
}

func (p Person) fullName() string {
	return p.first + " " + p.last
}

func main() {

	p1 := Singer{
		Person: Person{
			first: "Nina",
			last: "Simone",
			age: 33,
		},
		genre: "jazz",
	}
	p2 := Person{"Aretha", "Franklin", 25}
	p3 := Person{"Eartha", "Kitt", 28}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	fmt.Println(p3.first, p3.last, p3.age)
}