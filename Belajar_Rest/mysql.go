package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "xxx:yyy@tcp(zzz:3306)/database")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
