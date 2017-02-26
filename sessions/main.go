package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")

	if err == http.ErrNoCookie {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true // makes cookie available only on HTTPS
			HttpOnly: true, // cannot access cookie with Javascript
		}

		http.SetCookie(res, cookie)
	}

	fmt.Println(cookie)
}
