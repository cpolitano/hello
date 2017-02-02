package main

import (
	// "fmt"
	"net/http"
	"html/template"
	"log"
)

// the handler interface type
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

type thing int 

func (m thing) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var t thing
	http.ListenAndServe(":8080", t)
}

