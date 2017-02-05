package main

import (
	"io"
	"net/http"
)

// the handler interface type
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

func getBooks(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "books")
}

func getTapes(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "tapes")
}

func main() {
	http.HandleFunc("/books/", getBooks) // need trailing slash to handle /books/book_id 
	http.HandleFunc("/tapes", getTapes)

	http.ListenAndServe(":8080", nil)
}
