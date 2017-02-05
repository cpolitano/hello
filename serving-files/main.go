package main

import (
	"io"
	"net/http"
)

func hedgie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8" )
	io.WriteString(w, `<img src="/hedgie.jpg">`)
}

func getHedgie(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "hedgie.jpg")
}

func main() {
	http.HandleFunc("/", hedgie) 
	http.HandleFunc("/hedgie.jpg", getHedgie)
	http.ListenAndServe(":8080", nil) 
}
