// ===================================================================
// internal/database/queries.go (NEW FILE) - Driver-specific queries
// ===================================================================
package database

import "fmt"

// QueryBuilder builds queries for different database drivers
type QueryBuilder struct {
	driver DatabaseDriver
}

func NewQueryBuilder(driver DatabaseDriver) *QueryBuilder {
	return &QueryBuilder{driver: driver}
}

// Placeholder returns the correct parameter placeholder for the driver
func (qb *QueryBuilder) Placeholder(n int) string {
	switch qb.driver {
	case PostgreSQL:
		return fmt.Sprintf("$%d", n)
	case SQLite:
		return "?"
	default:
		return "?"
	}
}

// Now returns the current timestamp expression
func (qb *QueryBuilder) Now() string {
	switch qb.driver {
	case PostgreSQL:
		return "NOW()"
	case SQLite:
		return "datetime('now')"
	default:
		return "datetime('now')"
	}
}

// Serial returns the auto-increment column definition
func (qb *QueryBuilder) Serial() string {
	switch qb.driver {
	case PostgreSQL:
		return "SERIAL PRIMARY KEY"
	case SQLite:
		return "INTEGER PRIMARY KEY AUTOINCREMENT"
	default:
		return "INTEGER PRIMARY KEY AUTOINCREMENT"
	}
}

// TimestampType returns the timestamp column type
func (qb *QueryBuilder) TimestampType() string {
	switch qb.driver {
	case PostgreSQL:
		return "TIMESTAMP WITH TIME ZONE"
	case SQLite:
		return "DATETIME"
	default:
		return "DATETIME"
	}
}

// ===================================================================
// CATEGORY QUERIES
// ===================================================================

// BuildCreateUserQuery builds the create user query
func (qb *QueryBuilder) BuildCreateUserQuery() string {
	return fmt.Sprintf(`
    INSERT INTO users (username, email, password, created_at, updated_at) 
    VALUES (%s, %s, %s, %s, %s) 
    RETURNING id, username, email, created_at, updated_at`,
		qb.Placeholder(1), qb.Placeholder(2), qb.Placeholder(3),
		qb.Placeholder(4), qb.Placeholder(5))
}

// BuildGetUserQuery builds the get user query
func (qb *QueryBuilder) BuildGetUserQuery() string {
	return fmt.Sprintf(`
    SELECT id, username, email, password, created_at, updated_at 
    FROM users WHERE username = %s`, qb.Placeholder(1))
}

func (qb *QueryBuilder) BuildGetUserByIdQuery() string {
	return fmt.Sprintf(`
	SELECT id, username, email, created_at, updated_at FROM users WHERE id = %s`, qb.Placeholder(1))
}

// ===================================================================
// CATEGORY QUERIES
// ===================================================================

// BuildCreateTodoQuery builds the create todo query
func (qb *QueryBuilder) BuildCreateCategoryQuery() string {
	return fmt.Sprintf(`
    INSERT INTO categories (user_id, name, description, color, created_at, updated_at)
    VALUES (%s, %s, %s, %s, %s, %s) 
    RETURNING id, user_id, name, description, color, created_at, updated_at`,
		qb.Placeholder(1), qb.Placeholder(2), qb.Placeholder(3),
		qb.Placeholder(4), qb.Placeholder(5), qb.Placeholder(6))
}

// BuildGetUserQuery builds the get user query
func (qb *QueryBuilder) BuildGetCategoriesQuery() string {
	return fmt.Sprintf(`
    SELECT id, user_id, name, description, color, created_at, updated_at
    FROM categories WHERE user_id = %s
	ORDER BY name`,
		qb.Placeholder(1))
}

func (qb *QueryBuilder) BuildGetCategoriesByIdQuery() string {
	return fmt.Sprintf(`
    SELECT id, user_id, name, description, color, created_at, updated_at
    FROM categories WHERE id = %s AND user_id = %s`,
		qb.Placeholder(1), qb.Placeholder(2))
}

func (qb *QueryBuilder) BuildUpdateCategoriesQuery() string {
	return fmt.Sprintf(`
    UPDATE categories 
	SET name = %s, description = %s, color = %s, updated_at = %s 
	WHERE id = %s AND user_id = %s`,
		qb.Placeholder(1), qb.Placeholder(2), qb.Placeholder(3),
		qb.Placeholder(4), qb.Placeholder(5), qb.Placeholder(6))
}

func (qb *QueryBuilder) BuildDeleteCategoriesQuery() string {
	return fmt.Sprintf(`
    DELETE FROM categories 
	WHERE id = %s AND user_id = %s`,
		qb.Placeholder(1), qb.Placeholder(2))
}

// ===================================================================
// TODO QUERIES
// ===================================================================

// BuildCreateTodoQuery builds the create todo query
func (qb *QueryBuilder) BuildCreateTodoQuery() string {
	return fmt.Sprintf(`
    INSERT INTO todos (user_id, title, description, category_id, created_at, updated_at, frequency, priority) 
    VALUES (%s, %s, %s, %s, %s, %s, %s, %s) 
    RETURNING id, user_id, title, description, completed, category_id, created_at, updated_at, frequency, priority`,
		qb.Placeholder(1), qb.Placeholder(2), qb.Placeholder(3),
		qb.Placeholder(4), qb.Placeholder(5), qb.Placeholder(6),
		qb.Placeholder(7), qb.Placeholder(8))
}

// BuildGetTodoQuery
func (qb *QueryBuilder) BuildGetTodosQuery() string {
	return fmt.Sprintf(`
    SELECT 
        t.id, t.user_id, t.title, t.description, t.completed, t.category_id, t.created_at, t.updated_at, t.frequency, t.priority,
        c.id, c.user_id, c.name, c.description, c.color, c.created_at, c.updated_at
    FROM todos t 
    LEFT JOIN categories c ON t.category_id = c.id 
    WHERE t.user_id = %s
    ORDER BY t.created_at DESC`, qb.Placeholder(1))
}

// BuildGetTodoQuery
func (qb *QueryBuilder) BuildGetTodosByIdQuery() string {
	return fmt.Sprintf(`
    SELECT 
        t.id, t.user_id, t.title, t.description, t.completed, t.category_id, t.created_at, t.updated_at, t.frequency, t.priority,
        c.id, c.user_id, c.name, c.description, c.color, c.created_at, c.updated_at
    FROM todos t 
    LEFT JOIN categories c ON t.category_id = c.id 
    WHERE t.id = %s AND t.user_id = %s`,
		qb.Placeholder(1), qb.Placeholder(2))
}

func (qb *QueryBuilder) BuildUpdateTodoQuery() string {
	return fmt.Sprintf(`
	UPDATE todos 
	SET title = %s, description = %s, completed = %s, category_id = %s,
	updated_at = %s, frequency = %s, priority = %s 
	WHERE id = %s AND user_id = %s`,
		qb.Placeholder(1), qb.Placeholder(2), qb.Placeholder(3),
		qb.Placeholder(4), qb.Placeholder(5), qb.Placeholder(6),
		qb.Placeholder(7), qb.Placeholder(8), qb.Placeholder(9))
}

func (qb *QueryBuilder) BuildDeleteTodoQuery() string {
	return fmt.Sprintf(`
	DELETE FROM todos 
	WHERE id = %s AND user_id = %s`,
		qb.Placeholder(1), qb.Placeholder(2))
}

// ===================================================================
// MIGRATION QUERIES
// ===================================================================
func (qb *QueryBuilder) BuildRecordMigrationQuery() string {
	return fmt.Sprintf(`
    INSERT INTO schema_migrations (version, description, applied_at) 
    VALUES (%s, %s, %s) `,
		qb.Placeholder(1), qb.Placeholder(2), qb.Placeholder(3))
}
