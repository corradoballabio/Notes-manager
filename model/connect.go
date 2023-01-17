package model

import (
	"fmt"
	"database/sql"
	"log"
)

// db ref accessible throughout all the model package
var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:12345678!@tcp(localhost:3306)/mysql")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the db")

	con = db
	return db
}