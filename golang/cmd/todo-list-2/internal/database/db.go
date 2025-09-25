// internal/database/db.go (COMPLETELY UPDATED with Users)
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"todo-list-2/internal/logger"
	"todo-list-2/internal/migrations"
	"todo-list-2/internal/models"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type DB struct {
	conn    *sql.DB
	driver  DatabaseDriver
	queries *QueryBuilder
	logger  *logger.Logger
}

func New(dataSourceName string, basePath string, logger *logger.Logger) (*DB, error) {
	logger.Info("Creating Database from source", "dataSourceName", dataSourceName)
	conn, driver, err := Connect(dataSourceName, basePath)
	if err != nil {
		return nil, err
	}

	db := &DB{
		conn:    conn,
		driver:  driver,
		queries: NewQueryBuilder(driver),
		logger:  logger,
	}

	// Run migrations
	migrator := migrations.NewMigrator(conn, "sqlite", logger, basePath) //TODO
	if err := migrator.MigrateUp(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return db, nil
}

func NewFromConnection(conn *sql.DB, driver DatabaseDriver, logger *logger.Logger) *DB {
	return &DB{
		conn:    conn,
		driver:  driver,
		queries: NewQueryBuilder(driver),
		logger:  logger,
	}
}

func (db *DB) Close() error {
	return db.conn.Close()
}

// GetDriver returns the current database driver
func (db *DB) GetDriver() DatabaseDriver {
	return db.driver
}

// ===================================================================
// USER CRUD OPERATIONS
// ===================================================================

// Updated CreateUser method with driver-specific queries
func (db *DB) CreateUser(req models.RegisterRequest) (*models.User, error) {
	start := time.Now()

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		db.logger.LogDBOperation("hash_password", "users", 0, time.Since(start), err)
		return nil, err
	}

	// Check if username already exists
	var count int
	checkQuery := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE username = %s", db.queries.Placeholder(1))
	err = db.conn.QueryRow(checkQuery, req.Username).Scan(&count)
	if err != nil {
		db.logger.LogDBOperation("check_username", "users", 0, time.Since(start), err)
		return nil, err
	}
	if count > 0 {
		err := errors.New("username already exists")
		db.logger.LogDBOperation("check_username", "users", 0, time.Since(start), err)
		return nil, err
	}

	// Insert user
	query := db.queries.BuildCreateUserQuery()
	now := time.Now()
	row := db.conn.QueryRow(query, req.Username, req.Email, string(hashedPassword), now, now)

	user := &models.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		db.logger.LogDBOperation("create_user", "users", 0, time.Since(start), err)
		return nil, err
	}

	db.logger.LogDBOperation("create_user", "users", user.ID, time.Since(start), nil)
	db.logger.Info("New user created", "user_id", user.ID, "username", user.Username)

	// Create default categories for new user
	err = db.createDefaultCategoriesForUser(user.ID)
	if err != nil {
		db.logger.Warn("Failed to create default categories", "user_id", user.ID, "error", err.Error())
	} else {
		db.logger.Debug("Default categories created", "user_id", user.ID)
	}

	return user, nil
}

// Updated AuthenticateUser method
func (db *DB) AuthenticateUser(username, password string) (*models.User, error) {
	start := time.Now()

	query := db.queries.BuildGetUserQuery()

	user := &models.User{}
	row := db.conn.QueryRow(query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("user not found")
		}
		db.logger.LogDBOperation("authenticate_user", "users", 0, time.Since(start), err)
		return nil, err
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		err = errors.New("invalid password")
		db.logger.LogDBOperation("authenticate_user", "users", user.ID, time.Since(start), err)
		return nil, err
	}

	db.logger.LogDBOperation("authenticate_user", "users", user.ID, time.Since(start), nil)

	// Clear password before returning
	user.Password = ""
	return user, nil
}

func (db *DB) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?`

	user := &models.User{}
	row := db.conn.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DB) createDefaultCategoriesForUser(userID int) error {
	defaults := []struct {
		name, description, color string
	}{
		{"Personal", "Personal tasks and reminders", "#10B981"},
		{"Work", "Work-related todos", "#3B82F6"},
		{"Shopping", "Shopping lists and errands", "#F59E0B"},
		{"Health", "Health and fitness goals", "#EF4444"},
	}

	for _, cat := range defaults {
		now := time.Now()
		_, err := db.conn.Exec(
			"INSERT INTO categories (user_id, name, description, color, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
			userID, cat.name, cat.description, cat.color, now, now,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// ===================================================================
// CATEGORY CRUD OPERATIONS (UPDATED with user filtering)
// ===================================================================

func (db *DB) CreateCategory(userID int, req models.CreateCategoryRequest) (*models.Category, error) {
	query := `
    INSERT INTO categories (user_id, name, description, color, created_at, updated_at) 
    VALUES (?, ?, ?, ?, ?, ?) 
    RETURNING id, user_id, name, description, color, created_at, updated_at`

	now := time.Now()
	color := req.Color
	if color == "" {
		color = "#3B82F6"
	}

	row := db.conn.QueryRow(query, userID, req.Name, req.Description, color, now, now)

	category := &models.Category{}
	err := row.Scan(&category.ID, &category.UserID, &category.Name, &category.Description, &category.Color, &category.CreatedAt, &category.UpdatedAt)
	return category, err
}

func (db *DB) GetAllCategories(userID int) ([]models.Category, error) {
	query := `SELECT id, user_id, name, description, color, created_at, updated_at FROM categories WHERE user_id = ? ORDER BY name`

	rows, err := db.conn.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.UserID, &category.Name, &category.Description, &category.Color, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (db *DB) GetCategoryByID(userID, id int) (*models.Category, error) {
	query := `SELECT id, user_id, name, description, color, created_at, updated_at FROM categories WHERE id = ? AND user_id = ?`

	category := &models.Category{}
	row := db.conn.QueryRow(query, id, userID)
	err := row.Scan(&category.ID, &category.UserID, &category.Name, &category.Description, &category.Color, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (db *DB) UpdateCategory(userID, id int, req models.UpdateCategoryRequest) (*models.Category, error) {
	current, err := db.GetCategoryByID(userID, id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		current.Name = *req.Name
	}
	if req.Description != nil {
		current.Description = *req.Description
	}
	if req.Color != nil {
		current.Color = *req.Color
	}
	current.UpdatedAt = time.Now()

	query := `UPDATE categories SET name = ?, description = ?, color = ?, updated_at = ? WHERE id = ? AND user_id = ?`
	_, err = db.conn.Exec(query, current.Name, current.Description, current.Color, current.UpdatedAt, id, userID)
	if err != nil {
		return nil, err
	}

	return current, nil
}

func (db *DB) DeleteCategory(userID, id int) error {
	query := `DELETE FROM categories WHERE id = ? AND user_id = ?`
	_, err := db.conn.Exec(query, id, userID)
	return err
}

// ===================================================================
// TODO CRUD OPERATIONS (UPDATED with user filtering)
// ===================================================================

// Updated CreateTodo method
func (db *DB) CreateTodo(userID int, req models.CreateTodoRequest) (*models.Todo, error) {
	start := time.Now()

	// Verify category belongs to user if category_id is provided
	if req.CategoryID != nil {
		_, err := db.GetCategoryByID(userID, *req.CategoryID)
		if err != nil {
			return nil, errors.New("category not found or doesn't belong to user")
		}
	}

	query := db.queries.BuildCreateTodoQuery()
	now := time.Now()
	row := db.conn.QueryRow(query, userID, req.Title, req.Description, req.CategoryID, now, now)

	todo := &models.Todo{}
	err := row.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Completed, &todo.CategoryID, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		db.logger.LogDBOperation("create_todo", "todos", userID, time.Since(start), err)
		return nil, err
	}

	// Load category info if present
	if todo.CategoryID != nil {
		category, err := db.GetCategoryByID(userID, *todo.CategoryID)
		if err == nil {
			todo.Category = category
		}
	}

	db.logger.LogDBOperation("create_todo", "todos", userID, time.Since(start), nil)
	return todo, nil
}

func (db *DB) GetAllTodos(userID int) ([]models.Todo, error) {
	query := `
    SELECT 
        t.id, t.user_id, t.title, t.description, t.completed, t.category_id, t.created_at, t.updated_at,
        c.id, c.user_id, c.name, c.description, c.color, c.created_at, c.updated_at
    FROM todos t 
    LEFT JOIN categories c ON t.category_id = c.id 
    WHERE t.user_id = ?
    ORDER BY t.created_at DESC`

	rows, err := db.conn.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		var categoryID, catID, catUserID sql.NullInt64
		var catName, catDesc, catColor sql.NullString
		var catCreated, catUpdated sql.NullTime

		err := rows.Scan(
			&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Completed, &categoryID, &todo.CreatedAt, &todo.UpdatedAt,
			&catID, &catUserID, &catName, &catDesc, &catColor, &catCreated, &catUpdated,
		)
		if err != nil {
			return nil, err
		}

		if categoryID.Valid {
			id := int(categoryID.Int64)
			todo.CategoryID = &id
		}

		if catID.Valid {
			todo.Category = &models.Category{
				ID:          int(catID.Int64),
				UserID:      int(catUserID.Int64),
				Name:        catName.String,
				Description: catDesc.String,
				Color:       catColor.String,
				CreatedAt:   catCreated.Time,
				UpdatedAt:   catUpdated.Time,
			}
		}

		todos = append(todos, todo)
	}
	return todos, nil
}

func (db *DB) GetTodoByID(userID, id int) (*models.Todo, error) {
	query := `
    SELECT 
        t.id, t.user_id, t.title, t.description, t.completed, t.category_id, t.created_at, t.updated_at,
        c.id, c.user_id, c.name, c.description, c.color, c.created_at, c.updated_at
    FROM todos t 
    LEFT JOIN categories c ON t.category_id = c.id 
    WHERE t.id = ? AND t.user_id = ?`

	todo := &models.Todo{}
	var categoryID, catID, catUserID sql.NullInt64
	var catName, catDesc, catColor sql.NullString
	var catCreated, catUpdated sql.NullTime

	row := db.conn.QueryRow(query, id, userID)
	err := row.Scan(
		&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Completed, &categoryID, &todo.CreatedAt, &todo.UpdatedAt,
		&catID, &catUserID, &catName, &catDesc, &catColor, &catCreated, &catUpdated,
	)
	if err != nil {
		return nil, err
	}

	if categoryID.Valid {
		id := int(categoryID.Int64)
		todo.CategoryID = &id
	}

	if catID.Valid {
		todo.Category = &models.Category{
			ID:          int(catID.Int64),
			UserID:      int(catUserID.Int64),
			Name:        catName.String,
			Description: catDesc.String,
			Color:       catColor.String,
			CreatedAt:   catCreated.Time,
			UpdatedAt:   catUpdated.Time,
		}
	}

	return todo, nil
}

func (db *DB) UpdateTodo(userID, id int, req models.UpdateTodoRequest) (*models.Todo, error) {
	current, err := db.GetTodoByID(userID, id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		current.Title = *req.Title
	}
	if req.Description != nil {
		current.Description = *req.Description
	}
	if req.Completed != nil {
		current.Completed = *req.Completed
	}
	if req.CategoryID != nil {
		// Verify category belongs to user if not nil
		if *req.CategoryID != 0 {
			_, err := db.GetCategoryByID(userID, *req.CategoryID)
			if err != nil {
				return nil, errors.New("category not found or doesn't belong to user")
			}
		}
		current.CategoryID = req.CategoryID
	}
	current.UpdatedAt = time.Now()

	query := `UPDATE todos SET title = ?, description = ?, completed = ?, category_id = ?, updated_at = ? WHERE id = ? AND user_id = ?`
	_, err = db.conn.Exec(query, current.Title, current.Description, current.Completed, current.CategoryID, current.UpdatedAt, id, userID)
	if err != nil {
		return nil, err
	}

	return db.GetTodoByID(userID, id)
}

func (db *DB) DeleteTodo(userID, id int) error {
	query := `DELETE FROM todos WHERE id = ? AND user_id = ?`
	_, err := db.conn.Exec(query, id, userID)
	return err
}
