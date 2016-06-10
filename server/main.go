package main

import "net/http"
import "fmt"

func main() {
	fmt.Println("server running on port 8080")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)

	// _ is blank identifier, representing an error
	// res, _ := http.Get("http://dev.twitter.com")
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

