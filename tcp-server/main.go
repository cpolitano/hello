package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"strings"
)

func main() {
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
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// headers done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	// request
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]
	fmt.Println("METHOD", m)
	fmt.Println("URI", u)

	// multiplexer
	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/about" {
		about(conn)
	}
	if m == "GET" && u == "/contact" {
		contact(conn)
	}
	if m == "POST" && u == "/contact" {
		contactSubmit(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en"><head><meta charset="UTF-8"><title>Aloha</title></head>
	<body><h1>Aloha</h1><a href="/about">About</a><br><a href="/contact">Contact</a></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type text/html\r\n\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en"><head><meta charset="UTF-8"><title>Aloha</title></head>
	<body><h1>About</h1></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type text/html\r\n\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en"><head><meta charset="UTF-8"><title>Aloha</title></head>
	<body><h1>Contact</h1>
	<form method="POST" action="/contact"><input type="submit" value="contact"></form>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type text/html\r\n\r\n")
	fmt.Fprint(conn, body)
}

func contactSubmit(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en"><head><meta charset="UTF-8"><title>Aloha</title></head>
	<body><h1>Aloha</h1><h2>Contact submitted</h2>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type text/html\r\n\r\n")
	fmt.Fprint(conn, body)
}

