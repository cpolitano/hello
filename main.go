package main

import "net/http"

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    // println("\nhello!\n\n")
    w.Write([]byte("hello world"))
}