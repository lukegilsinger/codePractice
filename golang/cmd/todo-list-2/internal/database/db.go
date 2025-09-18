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
	return db, db.createTables()
}

func (db *DB) createTables() error {
	// Create categories table first (no dependencies)
	categoriesTable := `
    CREATE TABLE IF NOT EXISTS categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
        description TEXT,
        color TEXT DEFAULT '#3B82F6',
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )`

	if _, err := db.conn.Exec(categoriesTable); err != nil {
		return err
	}

	// Create todos table with foreign key to categories
	todosTable := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT,
        completed BOOLEAN DEFAULT FALSE,
        category_id INTEGER,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE SET NULL
    )`

	if _, err := db.conn.Exec(todosTable); err != nil {
		return err
	}

	// Insert default categories if table is empty
	return db.insertDefaultCategories()
}

func (db *DB) insertDefaultCategories() error {
	// Check if we have any categories
	var count int
	err := db.conn.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil // Categories already exist
	}

	// Insert default categories
	defaults := []struct {
		name, description, color string
	}{
		{"Personal", "Personal tasks and reminders", "#10B981"},
		{"Work", "Work-related todos", "#3B82F6"},
		{"Shopping", "Shopping lists and errands", "#F59E0B"},
		{"Health", "Health and fitness goals", "#EF4444"},
	}

	for _, cat := range defaults {
		_, err = db.conn.Exec(
			"INSERT INTO categories (name, description, color) VALUES (?, ?, ?)",
			cat.name, cat.description, cat.color,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

// CATEGORY CRUD OPERATIONS
func (db *DB) CreateCategory(req models.CreateCategoryRequest) (*models.Category, error) {
	query := `
    INSERT INTO categories (name, description, color, created_at, updated_at) 
    VALUES (?, ?, ?, ?, ?) 
    RETURNING id, name, description, color, created_at, updated_at`

	now := time.Now()
	color := req.Color
	if color == "" {
		color = "#3B82F6" // default blue
	}

	row := db.conn.QueryRow(query, req.Name, req.Description, color, now, now)

	category := &models.Category{}
	err := row.Scan(&category.ID, &category.Name, &category.Description, &category.Color, &category.CreatedAt, &category.UpdatedAt)
	return category, err
}

func (db *DB) GetAllCategories() ([]models.Category, error) {
	query := `SELECT id, name, description, color, created_at, updated_at FROM categories ORDER BY name`

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description, &category.Color, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (db *DB) GetCategoryByID(id int) (*models.Category, error) {
	query := `SELECT id, name, description, color, created_at, updated_at FROM categories WHERE id = ?`

	category := &models.Category{}
	row := db.conn.QueryRow(query, id)
	err := row.Scan(&category.ID, &category.Name, &category.Description, &category.Color, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (db *DB) UpdateCategory(id int, req models.UpdateCategoryRequest) (*models.Category, error) {
	current, err := db.GetCategoryByID(id)
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

	query := `UPDATE categories SET name = ?, description = ?, color = ?, updated_at = ? WHERE id = ?`
	_, err = db.conn.Exec(query, current.Name, current.Description, current.Color, current.UpdatedAt, id)
	if err != nil {
		return nil, err
	}

	return current, nil
}

func (db *DB) DeleteCategory(id int) error {
	query := `DELETE FROM categories WHERE id = ?`
	_, err := db.conn.Exec(query, id)
	return err
}

// UPDATED TODO CRUD OPERATIONS
func (db *DB) CreateTodo(req models.CreateTodoRequest) (*models.Todo, error) {
	query := `
    INSERT INTO todos (title, description, category_id, created_at, updated_at) 
    VALUES (?, ?, ?, ?, ?) 
    RETURNING id, title, description, completed, category_id, created_at, updated_at`

	now := time.Now()
	row := db.conn.QueryRow(query, req.Title, req.Description, req.CategoryID, now, now)

	todo := &models.Todo{}
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CategoryID, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Load category info if present
	if todo.CategoryID != nil {
		category, err := db.GetCategoryByID(*todo.CategoryID)
		if err == nil {
			todo.Category = category
		}
	}

	return todo, nil
}

func (db *DB) GetAllTodos() ([]models.Todo, error) {
	query := `
    SELECT 
        t.id, t.title, t.description, t.completed, t.category_id, t.created_at, t.updated_at,
        c.id, c.name, c.description, c.color, c.created_at, c.updated_at
    FROM todos t 
    LEFT JOIN categories c ON t.category_id = c.id 
    ORDER BY t.created_at DESC`

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		var categoryID, catID sql.NullInt64
		var catName, catDesc, catColor sql.NullString
		var catCreated, catUpdated sql.NullTime

		err := rows.Scan(
			&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &categoryID, &todo.CreatedAt, &todo.UpdatedAt,
			&catID, &catName, &catDesc, &catColor, &catCreated, &catUpdated,
		)
		if err != nil {
			return nil, err
		}

		if categoryID.Valid {
			id := int(categoryID.Int64)
			todo.CategoryID = &id
		}

		// Populate category if it exists
		if catID.Valid {
			todo.Category = &models.Category{
				ID:          int(catID.Int64),
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

func (db *DB) GetTodoByID(id int) (*models.Todo, error) {
	query := `
    SELECT 
        t.id, t.title, t.description, t.completed, t.category_id, t.created_at, t.updated_at,
        c.id, c.name, c.description, c.color, c.created_at, c.updated_at
    FROM todos t 
    LEFT JOIN categories c ON t.category_id = c.id 
    WHERE t.id = ?`

	todo := &models.Todo{}
	var categoryID, catID sql.NullInt64
	var catName, catDesc, catColor sql.NullString
	var catCreated, catUpdated sql.NullTime

	row := db.conn.QueryRow(query, id)
	err := row.Scan(
		&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &categoryID, &todo.CreatedAt, &todo.UpdatedAt,
		&catID, &catName, &catDesc, &catColor, &catCreated, &catUpdated,
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
			Name:        catName.String,
			Description: catDesc.String,
			Color:       catColor.String,
			CreatedAt:   catCreated.Time,
			UpdatedAt:   catUpdated.Time,
		}
	}

	return todo, nil
}

func (db *DB) UpdateTodo(id int, req models.UpdateTodoRequest) (*models.Todo, error) {
	current, err := db.GetTodoByID(id)
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
		current.CategoryID = req.CategoryID
	}
	current.UpdatedAt = time.Now()

	query := `UPDATE todos SET title = ?, description = ?, completed = ?, category_id = ?, updated_at = ? WHERE id = ?`
	_, err = db.conn.Exec(query, current.Title, current.Description, current.Completed, current.CategoryID, current.UpdatedAt, id)
	if err != nil {
		return nil, err
	}

	return db.GetTodoByID(id) // Return with populated category
}

func (db *DB) DeleteTodo(id int) error {
	query := `DELETE FROM todos WHERE id = ?`
	_, err := db.conn.Exec(query, id)
	return err
}
