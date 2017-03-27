package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "user:password@tcp(aws endpoint)/dbname?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)
	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.HandleFunc("/unicorns", unicorns)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "welcome to unicorn land")
}

func ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func instance(w http.ResponseWriter, req *http.Request) {
	s := getInstance()
	io.WriteString(w, s)
}

func unicorns(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT name FROM unicorns;`)
	check(err)

	msg := getInstance()
	msg += "\nRETRIEVED RECORDS:\n"
	var name string

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		msg += name + "\n"
	}

	fmt.Fprintln(w, msg)
}

func getInstance() string {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()
	return string(bs)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
