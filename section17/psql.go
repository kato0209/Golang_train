package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main2() {
	Db, _ = sql.Open("postgres", "user=postgres dbname=testdb password=p@ssword sslmode=disable")

	defer Db.Close()

}
