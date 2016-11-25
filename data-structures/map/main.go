package main

import (
	"fmt";
	"log";
	"net/http";
	"bufio"
)

func HashBucket(word string) int {
	return int(word[0])
}

func main() {
	// maps are for key/value storage
	// reference type pointing to underlying hash table / dictionary
	// value of uninitialized map is nil
	// NOT ordered
	m := make(map[string]int)
	m["f0"] = 1
	m["f1"] = 1
	m["f2"] = 2
	m["f3"] = 3
	m["f4"] = 5
	m["f5"] = 8

	delete(m, "f5")

	// will NOT print in order
	fmt.Println("map:", m)


	var greetings = map[string]string{} // composite literal
	greetings["Tulsi"] = "Aloha"
	greetings["Emilia"] = "Ciao"

	fmt.Println(greetings)

	salutations := map[int]string{
		0: "Hello",
		1: "Hola",
		2: "Aloha",
		3: "Ciao",
		4: "Salam", // need trailing comma
	}

	for key, val := range salutations {
		fmt.Println(key, "-", val)
	}	



	// creating a hash table "buckets" with words from Alice in Wonderland
	// goal is to create buckets that are as even as possible for most efficient lookup
	res, err := http.Get("http://www.gutenberg.org/files/11/11.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scans res text
	scanner := bufio.NewScanner(res.Body)
	// close stream at end of execution of main
	defer res.Body.Close()
	// split words
	scanner.Split(bufio.ScanWords)
	// create slice to hold counts
	buckets := make([]int, 200)
	// loop over words
	for scanner.Scan() {
		// divide into bucket based on int from first letter of word
		n := HashBucket(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:123])




}