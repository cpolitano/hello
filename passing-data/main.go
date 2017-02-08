package main

import (
	"io"
	"net/http"
	"log"
)

func getQueryParams(w http.ResponseWriter, req *http.Request) {
	value := req.FormValue("q")
	io.WriteString(w, "Do my search "+value)
}

func main() {
	http.HandleFunc("/", getQueryParams)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil)) 
}
