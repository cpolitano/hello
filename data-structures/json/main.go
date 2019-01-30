package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// capitalized struct name and field names
// mean they will be visible and usable outside package
// and inside main.go
type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

type Animal struct {
	Name  string
	Order string
}

func main() {

	// JSON MARSHAL
	// composite literal to instantiate a ColorGroup
	group1 := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Maroon", "Burgundy", "Oxblood", "Crimson"},
	}
	// json.Marshal turns the struct into json
	// Marshal takes any type and returns slice of bytes
	b, err := json.Marshal(group1)
	// always check for errors right after one might be returned
	// always check for errors.
	if err != nil {
		fmt.Println("err: ", err)
	}
	os.Stdout.Write(b)

	group2 := ColorGroup{
		ID:     2,
		Name:   "Yellows",
		Colors: []string{"Lemon", "Goldenrod", "Sunflower"},
	}

	groups := []ColorGroup{group1, group2}
	// returns byte slice
	bs, err := json.Marshal(groups)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(string(bs)) // convert byte slice to json string

	// JSON UNMARSHAL
	// insert raw json string literal into a slice of byte
	// json array of objects - convert to string, which is a slice of bytes
	var jsonBlob = []byte(`[{"Name": "Platypus", "Order": "Monotremata"}, {"Name": "Quoll", "Order": "Dasyuromor"}]`)

	// slice of animals
	var animals []Animal
	// unmarshal byte slice and store it at address of animals
	// Unmarshal takes slice of bytes and stores unmarshaled data in the location pointed to by &animals
	// Unmarshal ONLY returns an error if anything
	err = json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Printf("%+v", animals)
}
