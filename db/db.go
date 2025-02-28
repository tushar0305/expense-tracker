package db

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createTableUser := `
	CREATE TABLE IF NOT EXISTS USERS(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`

	_, err := Db.Exec(createTableUser)
	if err != nil {
		log.Fatalf("Could not create user table: %v", err)
	}
}