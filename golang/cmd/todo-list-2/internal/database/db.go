package database

import (
	"database/sql"
	"time"
	"todo-list-2/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	conn *sql.DB
}

func New(dataSourceName string) (*DB, error) {
	conn, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	db := &DB{conn: conn}
	return db, db.createTable()
}

func (db *DB) createTable() error {
	query := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT,
        completed BOOLEAN DEFAULT FALSE,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )`

	_, err := db.conn.Exec(query)
	return err
}

func (db *DB) Close() error {
	return db.conn.Close()
}

// CRUD Operations
func (db *DB) CreateTodo(req models.CreateTodoRequest) (*models.Todo, error) {
	query := `
    INSERT INTO todos (title, description, created_at, updated_at) 
    VALUES (?, ?, ?, ?) 
    RETURNING id, title, description, completed, created_at, updated_at`

	now := time.Now()
	row := db.conn.QueryRow(query, req.Title, req.Description, now, now)

	todo := &models.Todo{}
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	return todo, err
}

func (db *DB) GetAllTodos() ([]models.Todo, error) {
	query := `SELECT id, title, description, completed, created_at, updated_at FROM todos ORDER BY created_at DESC`

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (db *DB) GetTodoByID(id int) (*models.Todo, error) {
	query := `SELECT id, title, description, completed, created_at, updated_at FROM todos WHERE id = ?`

	todo := &models.Todo{}
	row := db.conn.QueryRow(query, id)
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (db *DB) UpdateTodo(id int, req models.UpdateTodoRequest) (*models.Todo, error) {
	// Get current todo
	current, err := db.GetTodoByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Title != nil {
		current.Title = *req.Title
	}
	if req.Description != nil {
		current.Description = *req.Description
	}
	if req.Completed != nil {
		current.Completed = *req.Completed
	}
	current.UpdatedAt = time.Now()

	query := `UPDATE todos SET title = ?, description = ?, completed = ?, updated_at = ? WHERE id = ?`
	_, err = db.conn.Exec(query, current.Title, current.Description, current.Completed, current.UpdatedAt, id)
	if err != nil {
		return nil, err
	}

	return current, nil
}

func (db *DB) DeleteTodo(id int) error {
	query := `DELETE FROM todos WHERE id = ?`
	_, err := db.conn.Exec(query, id)
	return err
}
