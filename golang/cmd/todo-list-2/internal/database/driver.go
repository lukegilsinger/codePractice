// internal/database/driver.go (NEW FILE) - Database driver abstraction
package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"           // PostgreSQL driver
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// DatabaseDriver represents different database types
type DatabaseDriver int

const (
	SQLite DatabaseDriver = iota
	PostgreSQL
)

// DriverConfig holds driver-specific configuration
type DriverConfig struct {
	Driver     DatabaseDriver
	DataSource string
}

// ParseDataSource determines database driver from connection string
func ParseDataSource(dataSource string) *DriverConfig {
	if strings.HasPrefix(dataSource, "postgres://") ||
		strings.HasPrefix(dataSource, "postgresql://") {
		return &DriverConfig{
			Driver:     PostgreSQL,
			DataSource: dataSource,
		}
	}

	// Default to SQLite for file paths or :memory:
	return &DriverConfig{
		Driver:     SQLite,
		DataSource: dataSource,
	}
}

// Connect creates a database connection based on the data source
func Connect(dataSource string) (*sql.DB, DatabaseDriver, error) {
	config := ParseDataSource(dataSource)

	var driverName string
	switch config.Driver {
	case PostgreSQL:
		driverName = "postgres"
	case SQLite:
		driverName = "sqlite3"
	default:
		return nil, SQLite, fmt.Errorf("unsupported database driver")
	}
	fmt.Println("SOURCE: ", config.DataSource)
	conn, err := sql.Open(driverName, config.DataSource)
	if err != nil {
		return nil, config.Driver, err
	}

	// Test the connection
	if err := conn.Ping(); err != nil {
		conn.Close()
		return nil, config.Driver, fmt.Errorf("failed to ping database: %w", err)
	}

	return conn, config.Driver, nil
}
