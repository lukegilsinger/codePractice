package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB initializes the database connection and creates the tasks table if it doesn't exist.
func InitDB(databaseFile string) {
	fmt.Println("initializing sqlite db from " + databaseFile)
	var err error

	db, err = sql.Open("sqlite3", "./"+databaseFile)
	if err != nil {
		panic(err)
	}

	// Create tasks table if it doesn't exist
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT,
        status TEXT NOT NULL,
        due_date DATE
    );`
	if _, err = db.Exec(sqlStmt); err != nil {
		panic(err)
	}
	fmt.Println("sqlite db is running")
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}
