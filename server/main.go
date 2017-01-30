package main

import (
	"fmt"
	// "io"
	"log"
	"net"
	"net/http"
	"bufio"
)

func main() {
	// fmt.Println("server running on port 8080")
	// http.HandleFunc("/", hello)
	// http.ListenAndServe(":8080", nil)

	// connect through command line with "telnet localhost 8080"
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

		go handle(conn)

		// io.WriteString(conn, "\nHello from TCP server\n")
		// fmt.Fprintln(conn, "How are you doing?")
		// fmt.Fprintf(conn, "%v", "Well, I hope!\n")

		// conn.Close()
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

