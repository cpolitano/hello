package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "oh hey, running on AWS")
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}
