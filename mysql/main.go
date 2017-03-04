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
	// connect to db
	// user:password@tcp(aws endpoint)/dbname?charset=utf8
	db, err = sql.Open("mysql", "user:password@tcp(aws endpoint)/dbname?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "completed connection")
	check(err)
}

func create(w http.ResponseWriter, req *http.Request) {

	statement, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)

	r, err := statement.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func insert(w http.ResponseWriter, req *http.Request) {

	statement, err := db.Prepare(`INSERT INTO customer VALUES ("BB");`)
	check(err)

	r, err := statement.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
