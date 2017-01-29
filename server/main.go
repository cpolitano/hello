package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func main() {
	// fmt.Println("server running on port 8080")
	// http.HandleFunc("/", hello)
	// http.ListenAndServe(":8080", nil)

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	} 
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How are you doing?")
		fmt.Fprintf(conn, "%v", "Well, I hope!\n")

		conn.Close()
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

