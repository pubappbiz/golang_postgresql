package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func dbConnectionAndClose() {
	db, err := sql.Open("postgres", "dbname=golang_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	dbConnectionAndClose()
	fmt.Println("SUCCESS!")
}