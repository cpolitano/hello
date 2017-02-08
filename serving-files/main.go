package main

import (
	"io"
	"net/http"
	"log"
)

func hedgie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8" )
	io.WriteString(w, `<img src="/resources/hedgie.jpg">`)
}

func main() {
	// http.Handle("/", http.FileServer(http.Dir("."))) // serve all files in .
	// http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets")))) // serve all files in assets
	// http.HandleFunc("/hedgie", hedgie)

	// if index.html is present, go will automatically serve it
	// log.Fatal will catch any error returned from ListenAndServe
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))) 
}
