// ===================================================================
// cmd/migrate/main.go (NEW FILE) - CLI tool for migrations
// ===================================================================
package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"todo-list-2/internal/config"

	"todo-list-2/internal/logger"
	"todo-list-2/internal/migrations"

	_ "github.com/lib/pq"           // PostgreSQL driver
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Initialize logger
	logger := logger.Init("info", "text")

	// Load config
	cfg := config.Load()

	// Connect to database

	var driverName string
	var connString string
	switch cfg.DatabaseType {
	case "postgres":
		fmt.Println("HERE")
		driverName = "postgres"
		connString = cfg.DatabaseURL
	case "sqlite":

		fmt.Println("HEREs")
		driverName = "sqlite3"
		connString = fmt.Sprintf("%s/%s", cfg.BasePath, cfg.DatabaseURL)
	default:

		fmt.Println("HEREo")
		driverName = "sqlite3"
		connString = fmt.Sprintf("%s/%s", cfg.BasePath, cfg.DatabaseURL)
	}
	fmt.Println("TYPE:", cfg.DatabaseType)
	fmt.Println("DRIVER: ", driverName)
	fmt.Println("SOURCE: ", connString)

	db, err := sql.Open(driverName, connString)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	migrator := migrations.NewMigrator(db, cfg.DatabaseType, logger, cfg.BasePath)

	command := os.Args[1]

	switch command {
	case "up":
		if err := migrator.MigrateUp(); err != nil {
			logger.Error("Migration failed", "error", err.Error())
			os.Exit(1)
		}

	case "down":
		steps := 1
		if len(os.Args) > 2 {
			if s, err := strconv.Atoi(os.Args[2]); err == nil {
				steps = s
			}
		}
		if err := migrator.MigrateDown(steps); err != nil {
			logger.Error("Rollback failed", "error", err.Error())
			os.Exit(1)
		}

	case "status":
		if err := migrator.Status(); err != nil {
			logger.Error("Status check failed", "error", err.Error())
			os.Exit(1)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  go run cmd/migrate/main.go up              # Apply all pending migrations")
	fmt.Println("  go run cmd/migrate/main.go down [steps]    # Rollback migrations (default: 1 step)")
	fmt.Println("  go run cmd/migrate/main.go status          # Show migration status")
}
