package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
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

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var t thing
	http.ListenAndServe(":8080", t)
}
