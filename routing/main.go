package main

import (
	"io"
	"net/http"
	"html/template"
)

// the handler interface type
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

func getIndex(res http.ResponseWriter, req *http.Request) {
	data := struct {
		Title        string
	}{
		"Welcome",
	}
	tpl.ExecuteTemplate(res, "index.gohtml", data)
}

func getDog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog")
}

func getProfile(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "me")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(getIndex)) 
	http.Handle("/dog/", http.HandlerFunc(getDog)) // need trailing slash to handle /dog/dog_id 
	http.Handle("/me", http.HandlerFunc(getProfile))

	http.ListenAndServe(":8080", nil) // nil indicates use of default mux
}
