package main

import (
	"io"
	"net/http"
	"log"
	"io/ioutil"
	"os"
	"path/filepath"
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

// accepts file and saves to server
func handleContact(w http.ResponseWriter, req *http.Request) {
	var s string
	
	if req.Method == http.MethodPost {
		// get file
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} 
		defer f.Close()

		// read file as byte slice
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} 
		s = string(bs)

		// save file
		destination, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} 
		defer destination.Close()

		_, err = destination.Write(bs) // writes bytes to file
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} 
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="POST" enctype="multipart/form-data">
		 <input type="file" name="q">
		 <input type="submit">
		</form>
		<br>`+s)
}

func main() {
	http.HandleFunc("/", getQueryParams)
	http.HandleFunc("/contact", handleContact)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil)) 
}
