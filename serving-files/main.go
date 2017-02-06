package main

import (
	"io"
	"net/http"
)

func hedgie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8" )
	io.WriteString(w, `<img src="/hedgie.jpg">`)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("."))) // serve all files in .
	http.HandleFunc("/hedgie", hedgie)
	http.ListenAndServe(":8080", nil) 
}
