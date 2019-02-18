package main

import (
	"fmt"
	"sort"
)

type person struct {
	first     string
	age       int
	vegetable string
}

type ByAge []person

// Sort Interface is type interface and takes methods Len, Less, and Swap
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

type ByVegetable []person

func (v ByVegetable) Len() int           { return len(v) }
func (v ByVegetable) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v ByVegetable) Less(i, j int) bool { return v[i].vegetable < v[j].vegetable }

func main() {
	xi := []int{5, 13, 8, 3, 75, 34, 7, 2, 52, 25, 26}
	xs := []string{"Mango", "Papaya", "Aguacate", "Chayote", "Zapote"}
	xf := []float64{5.098, 2.019, 4.00003, 1.9, 1.922, 3.245667}

	sort.Ints(xi) // no reassignment needed, changes slice in place
	// a SLICE is a pointer, length, and capacity

	sort.Strings(xs)
	sort.Float64s(xf)

	fmt.Println(xi)
	fmt.Println(xs)
	fmt.Println(xf)

	// sorting slices and types
	p1 := person{"James", 33, "beet"}
	p2 := person{"Pat", 28, "beet"}
	p3 := person{"Clare", 32, "kale"}
	p4 := person{"Sara", 30, "broccoli"}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)

	// conversion - convert people into type ByAge
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByVegetable(people))
	fmt.Println(people)
}
