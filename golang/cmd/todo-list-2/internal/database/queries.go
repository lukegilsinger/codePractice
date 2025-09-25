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

// BuildCreateTodoQuery builds the create todo query
func (qb *QueryBuilder) BuildCreateTodoQuery() string {
	return fmt.Sprintf(`
    INSERT INTO todos (user_id, title, description, category_id, created_at, updated_at) 
    VALUES (%s, %s, %s, %s, %s, %s) 
    RETURNING id, user_id, title, description, completed, category_id, created_at, updated_at`,
		qb.Placeholder(1), qb.Placeholder(2), qb.Placeholder(3),
		qb.Placeholder(4), qb.Placeholder(5), qb.Placeholder(6))
}
