package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

func main1() {
	DbConnection, _ := sql.Open("sqlite3", "./section17/example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS persons(
                name STRING,
				age  INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

}
