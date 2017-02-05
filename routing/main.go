package main

import (
	"io"
	"net/http"
)

// the handler interface type
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

type getBooks int
type getTapes int

func (b getBooks) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "books")
}

func (t getTapes) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "tapes")
}

func main() {
	var b getBooks
	var t getTapes

	mux := http.NewServeMux()
	mux.Handle("/books", b)
	mux.Handle("/tapes", t)

	http.ListenAndServe(":8080", mux)
}
