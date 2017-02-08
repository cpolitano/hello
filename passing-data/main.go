package main

import (
	"io"
	"net/http"
	"log"
)

func getQueryParams(w http.ResponseWriter, req *http.Request) {
	value := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="GET">
		 <input type="text" name="q">
		 <input type="submit">
		</form>
		<br>`+value)
	// POST submits data in body, GET submits data in query params
}

func main() {
	http.HandleFunc("/", getQueryParams)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil)) 
}
