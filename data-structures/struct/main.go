package main

import (
	"fmt"
	"encoding/json"
)

// in Go we create types, not classes
// a struct is a sequence of named elements called fields
// fields have a name and a type
// capitalized props are "exported", lowercase are "unexported"
// tag tells struct to accept a value with a different key

type Person struct {
	First string
	Last string
	Age int `json:"Years"` // tag
}

type Singer struct {
	Person
	Genre string
}

func (p Person) fullName() string {
	return p.First + " " + p.Last
}

func main() {

	p1 := Singer{
		Person: Person{
			First: "Nina",
			Last: "Simone",
			Age: 33,
		},
		Genre: "jazz",
	}
	fmt.Println(p1.fullName())

	// Marshalling
	p2 := Person{"Aretha", "Franklin", 25}
	// turn struct data into json
	bs, _ := json.Marshal(p2)
	fmt.Println(string(bs))

	// Unmarshalling
	var p3 Person // initializes at zero value with empty strings
	// creates a slice of bytes with a string literal with json
	bs2 := []byte(`{"First": "Eartha", "Last": "Kitt", "Years": 28}`)
	json.Unmarshal(bs2, &p3)
	fmt.Println(p3)

}