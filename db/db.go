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
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	_, err := Db.Exec(createTableUser)
	if err != nil {
		log.Fatalf("Could not create user table: %v", err)
	}

	expenseTable := `
	CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		userId INTEGER NOT NULL,
		amount REAL NOT NULL,
		category TEXT NOT NULL,
		date DATETIME NOT NULL,
		description TEXT NOT NULL,
		FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
	);`

	_, err = Db.Exec(expenseTable)
	if err != nil {
		log.Fatalf("Could not create expenses table: %v", err)
	}
}
