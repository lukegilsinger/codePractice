package db

import (
	"database/sql"
	"fmt"
	"os"

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

	createScripts := []string{
		"create_tasks",
		"create_items",
		"create_todo_tasks",
		"create_task_history",
	}
	for _, fileName := range createScripts {
		createTable(fileName)
	}

	fmt.Println("sqlite db is running")
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}

func createTable(sqlFile string) {
	// Read the SQL file
	filePath := fmt.Sprintf("./db/%v.sql", sqlFile)
	fmt.Printf("creating table from %s\n", filePath)
	sqlStmt, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read SQL file: %v\n", err)
	}
	if _, err = db.Exec(string(sqlStmt)); err != nil {
		panic(err)
	}
}
