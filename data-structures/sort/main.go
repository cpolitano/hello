package main

import (
  "fmt"
  "sort"
)

func main() {
  xi := []int{5, 13, 8, 3, 75, 34, 7, 2, 52, 25, 26}
  xs := []string{"Mango", "Papaya", "Aguacate", "Chayote", "Zapote"}

  sort.Ints(xi) // no reassignment needed, changes slice in place
  // a SLICE is a pointer, length, and capacity

  sort.Strings(xs)

  fmt.Println(xi)
  fmt.Println(xs)
}
