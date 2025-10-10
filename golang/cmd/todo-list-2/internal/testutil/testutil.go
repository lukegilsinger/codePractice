// internal/testutil/testutil.go (NEW FILE) - Test utilities
package testutil

import (
	"database/sql"
	"testing"
	"todo-list-2/internal/logger"
	"todo-list-2/internal/migrations"
	"todo-list-2/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

// TestDB wraps sql.DB for testing (avoiding import cycle)
type TestDB struct {
	Conn   *sql.DB
	logger *logger.Logger
}

// SetupTestDB creates an in-memory SQLite database for testing
func SetupTestDB(t *testing.T) *TestDB {
	// Initialize logger
	logger := logger.Init("info", "text")
	logger.LogStartup("8080")
	// Use in-memory SQLite database
	conn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	// Run migrations to set up schema
	migrator := migrations.NewMigrator(conn, "sqlite", logger, "./")
	if err := migrator.MigrateUp(); err != nil {
		t.Fatalf("Failed to run migrations: %v", err)
	}

	return &TestDB{conn, logger}
}

// CreateTestUser creates a test user and returns it
func CreateTestUser(t *testing.T, db *TestDB) *models.User {
	// Directly insert user to avoid import cycle
	hashedPassword := "$2a$10$N9qo8uLOickgx2ZMRZoMye6xhDdvU9xLKzVe5e8d1LVz9Z.4U0D.G" // "password123" hashed

	query := `
    INSERT INTO users (username, email, password, created_at, updated_at) 
    VALUES (?, ?, ?, datetime('now'), datetime('now'))`

	result, err := db.Conn.Exec(query, "testuser", "test@example.com", hashedPassword)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	userID, _ := result.LastInsertId()

	return &models.User{
		ID:       int(userID),
		Username: "testuser",
		Email:    "test@example.com",
	}
}

// CreateTestCategory creates a test category for a user
func CreateTestCategory(t *testing.T, db *TestDB, userID int) *models.Category {
	query := `
    INSERT INTO categories (user_id, name, description, color, created_at, updated_at)
    VALUES (?, ?, ?, ?, datetime('now'), datetime('now'))`

	result, err := db.Conn.Exec(query, userID, "Test Category", "A test category", "#FF0000")
	if err != nil {
		t.Fatalf("Failed to create test category: %v", err)
	}

	categoryID, _ := result.LastInsertId()

	return &models.Category{
		ID:          int(categoryID),
		UserID:      userID,
		Name:        "Test Category",
		Description: "A test category",
		Color:       "#FF0000",
	}
}

// CreateTestTodo creates a test todo for a user
func CreateTestTodo(t *testing.T, db *TestDB, userID int, categoryID *int) *models.Todo {
	query := `
    INSERT INTO todos (user_id, title, description, category_id, created_at, updated_at)
    VALUES (?, ?, ?, ?, datetime('now'), datetime('now'))`

	result, err := db.Conn.Exec(query, userID, "Test Todo", "A test todo", categoryID)
	if err != nil {
		t.Fatalf("Failed to create test todo: %v", err)
	}

	todoID, _ := result.LastInsertId()

	return &models.Todo{
		ID:          int(todoID),
		UserID:      userID,
		Title:       "Test Todo",
		Description: "A test todo",
		CategoryID:  categoryID,
	}
}
