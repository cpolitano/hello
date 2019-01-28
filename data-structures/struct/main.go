package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// in Go we create types, not classes
// a struct is a sequence of named elements called fields
// fields have a name and a type
// capitalized props are "exported", lowercase are "unexported"
// tag tells struct to accept a value with a different key

// a struct is a composite data type. AKA an aggregate data type
// structs allow us to compose together values of different types

type Person struct {
	First string
	Last  string
	// could also declare as First, Last string
	Age   int `json:"Years"` // tag
}

type Singer struct {
	Person // anonymous or embedded field
	Genre string // identifier, type
}

// Method Set
// set of methods attached to type
func (p Person) fullName() string {
	return p.First + " " + p.Last
}

func main() {

	// composite literal
	// perferred way to "create value of type Singer"
	// we don't say object or class in Golang :)
	p1 := Singer{
		Person: Person{ // must define type Person inside type Singer
			First: "Nina",
			Last:  "Simone",
			Age:   33,
		},
		Genre: "jazz", // must include trailing commas
	}

	fmt.Println(p1)
	fmt.Println(p1.fullName())
	// access fields with dot notation
	// values of inner type Person get "promoted" to outer type Singer
	// so First and Genre are both accessible from p1
	fmt.Println(p1.First, p1.Genre)

	// Marshalling - converts struct into json key values
	p2 := Person{"Aretha", "Franklin", 25}
	// turn struct data into json
	bs, _ := json.Marshal(p2)
	fmt.Println(string(bs))

	// Unmarshalling
	var p3 Person // initializes at zero value with empty strings
	// creates a slice of bytes with a string literal with json
	bs2 := []byte(`{"First": "Eartha", "Last": "Kitt", "Years": 28}`)
	// turns slice of bytes into Person struct, save at &p3, pointer to initialized p3 Person
	json.Unmarshal(bs2, &p3) // needs address/pointer
	fmt.Println(p3)

	// Encoding - Writer
	// same result as marshalling struct to json
	json.NewEncoder(os.Stdout).Encode(p1)

	// Decoding - Reader
	var p4 Person
	rdr := strings.NewReader(`{"First": "Amy", "Last": "Winehouse", "Years": 27}`)
	// decode json and send to struct address &p4
	json.NewDecoder(rdr).Decode(&p4) // needs address/pointer
	fmt.Println(p4)
}
