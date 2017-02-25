package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visit-counter")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "visit-counter",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}

	if count > 3 {
		cookie.MaxAge = -1
	} else {
		count++
		cookie.Value = strconv.Itoa(count)
	}

	http.SetCookie(res, cookie)
	io.WriteString(res, cookie.Value)
}
