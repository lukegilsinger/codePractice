package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// var db *sql.DB // Global variable for the da÷tabase connection

func setupTestDB() func() {
	var err error
	db, err = sql.Open("sqlite3", ":memory:") // Use an in-memory database for testing
	if err != nil {
		panic(err)
	}

	// Create the tasks table
	sqlStmt := `
    CREATE TABLE tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT,
        status TEXT NOT NULL,
        due_date DATE
    );`
	if _, err = db.Exec(sqlStmt); err != nil {
		panic(err)
	}

	// Return a function to close the database connection
	return func() {
		db.Close()
	}
}

func TestCreateTask(t *testing.T) {
	teardown := setupTestDB()
	defer teardown()

	task := Task{
		Title:       "Test Task",
		Description: "This is a test task.",
		Status:      "Pending",
		DueDate:     "2025-12-31",
	}
	taskJSON, _ := json.Marshal(task)

	req, err := http.NewRequest("POST", "/add-task", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createTaskHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestGetTasks(t *testing.T) {
	teardown := setupTestDB()
	defer teardown()

	// Insert a test task
	_, err := db.Exec("INSERT INTO tasks (title, description, status, due_date) VALUES (?, ?, ?, ?)",
		"Test Task", "This is a test task.", "Pending", "2025-12-31")
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/get-tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTasksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var tasks []Task
	err = json.Unmarshal(rr.Body.Bytes(), &tasks)
	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(tasks))
	}
}

func TestUpdateTask(t *testing.T) {
	teardown := setupTestDB()
	defer teardown()

	// Insert a test task
	_, err := db.Exec("INSERT INTO tasks (title, description, status, due_date) VALUES (?, ?, ?, ?)",
		"Test Task", "This is a test task.", "Pending", "2025-12-31")
	if err != nil {
		t.Fatal(err)
	}

	// Update the task
	updatedTask := Task{
		ID:          1,
		Title:       "Updated Task",
		Description: "This is an updated task.",
		Status:      "In Progress",
		DueDate:     "2025-11-30",
	}
	taskJSON, _ := json.Marshal(updatedTask)

	req, err := http.NewRequest("PUT", "/update-task", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateTaskHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	// Verify the update
	row := db.QueryRow("SELECT title, status FROM tasks WHERE id = ?", 1)
	var title, status string
	err = row.Scan(&title, &status)
	if err != nil {
		t.Fatal(err)
	}

	if title != "Updated Task" || status != "In Progress" {
		t.Errorf("task was not updated correctly: got title %v, status %v", title, status)
	}
}

func TestDeleteTask(t *testing.T) {
	teardown := setupTestDB()
	defer teardown()

	// Insert a test task
	_, err := db.Exec("INSERT INTO tasks (title, description, status, due_date) VALUES (?, ?, ?, ?)",
		"Test Task", "This is a test task.", "Pending", "2025-12-31")
	if err != nil {
		t.Fatal(err)
	}

	// Delete the task
	req, err := http.NewRequest("DELETE", "/delete-task?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteTaskHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	// Verify the task was deleted
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM tasks WHERE id = ?", 1).Scan(&count)
	if err != nil {
		t.Fatal(err)
	}

	if count != 0 {
		t.Errorf("task was not deleted, expected count 0, got %d", count)
	}
}
