package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
	"strconv"
)

var (
	id int
	languageName string
	development string
)

func dbCRUD() {
	// Open / Close
	db, err := sql.Open("postgres", "dbname=golang_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Send ping
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Begin transaction / Rollback
	tran, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tran.Commit()

	// Create
	sqlCreate, err := db.Prepare("insert into programing_language (language_name, development, created_at, updated_at) values ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	response := sqlCreate.QueryRow("Go", "Google", time.Now(), time.Now())
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	var insertId int64
	_ = response.Scan(&insertId)

	fmt.Println("INSERT id: " + strconv.FormatInt(insertId, 10))

	// Read
	sqlRead, err := db.Prepare("select id, language_name, development from programing_language")
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	rows, err := sqlRead.Query()
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &languageName, &development)
		if err != nil {
			tran.Rollback()
			log.Fatal(err)
		}
		fmt.Println("SELECT id: " + strconv.Itoa(id) + " name: " + languageName + " develop: " + development)
	}

	// Update
	sqlUpdate, err := db.Prepare("update programing_language set language_name = $1, development = $2, updated_at = $3 where id = $4")
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	_, err = sqlUpdate.Query("PHP", "The PHP Group", time.Now(), insertId)
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	fmt.Println("UPDATED!")

	// Delete
	sqlDelete, err := db.Prepare("delete from programing_language where id = $1")
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	_, err = sqlDelete.Query(insertId)
	if err != nil {
		tran.Rollback()
		log.Fatal(err)
	}
	fmt.Println("DELETED!")
}

func main() {
	dbCRUD()
	fmt.Println("SUCCESS!")
}
