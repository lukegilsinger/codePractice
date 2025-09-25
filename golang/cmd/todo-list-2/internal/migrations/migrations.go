// ===================================================================
// internal/migrations/migrations.go (UPDATED) - Driver-aware migrations
// ===================================================================
package migrations

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"todo-list-2/internal/logger"
)

type Migration struct {
	Version     int
	Description string
	UpSQL       string
	DownSQL     string
}

type Migrator struct {
	db            *sql.DB
	driver        string
	migrationsDir string
	logger        *logger.Logger
}

func NewMigrator(db *sql.DB, driver string, logger *logger.Logger, basePath string) *Migrator {
	return &Migrator{
		db:            db,
		driver:        driver,
		migrationsDir: fmt.Sprintf("%s/%s", basePath, "internal/migrations"),
		logger:        logger,
	}
}

// LoadMigrations reads migration files, choosing driver-specific files when available
func (m *Migrator) LoadMigrations() ([]Migration, error) {
	migrations := make(map[int]*Migration)

	files, err := os.ReadDir(m.migrationsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read migrations directory: %w", err)
	}

	// Pattern: 001_description_up.sql or 001_description_down.sql
	pattern := regexp.MustCompile(`^(\d{3})_(.+)_(up|down)\.sql`)
	// // Pattern: 001_description_up.sql or 001_description_postgres_up.sql
	// pattern := regexp.MustCompile(`^(\d{3})_(.+?)(?:_(postgres|sqlite))?_(up|down)\.sql$`)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		matches := pattern.FindStringSubmatch(file.Name())
		if len(matches) < 4 {
			continue
		}

		versionStr := matches[1]
		description := strings.ReplaceAll(matches[2], "_", " ")
		// driverSpecific := matches[3]
		direction := matches[3]

		version, err := strconv.Atoi(versionStr)
		if err != nil {
			continue
		}

		// Skip driver-specific files that don't match current driver
		// if driverSpecific != "" {
		// 	continue // TODO
		// }

		// Read SQL content
		content, err := os.ReadFile(filepath.Join(m.migrationsDir, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("failed to read migration file %s: %w", file.Name(), err)
		}

		sqlContent := strings.TrimSpace(string(content))
		if sqlContent == "" {
			continue
		}

		// Get or create migration entry
		if migrations[version] == nil {
			migrations[version] = &Migration{
				Version:     version,
				Description: description,
			}
		}

		// Set SQL based on direction
		if direction == "up" {
			migrations[version].UpSQL = sqlContent
		} else if direction == "down" {
			migrations[version].DownSQL = sqlContent
		}
	}

	// Convert map to slice and sort by version
	var result []Migration
	for _, migration := range migrations {
		if migration.UpSQL != "" {
			result = append(result, *migration)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Version < result[j].Version
	})

	return result, nil
}

// createMigrationTable ensures the migrations tracking table exists
func (m *Migrator) createMigrationTable() error {
	query := `
        CREATE TABLE IF NOT EXISTS schema_migrations (
            version INTEGER PRIMARY KEY,
            description TEXT NOT NULL,
            applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `
	_, err := m.db.Exec(query)
	return err
}

// GetAppliedMigrations returns which migrations have already been applied
func (m *Migrator) GetAppliedMigrations() ([]int, error) {
	if err := m.createMigrationTable(); err != nil {
		return nil, err
	}

	rows, err := m.db.Query("SELECT version FROM schema_migrations ORDER BY version")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var versions []int
	for rows.Next() {
		var version int
		if err := rows.Scan(&version); err != nil {
			return nil, err
		}
		versions = append(versions, version)
	}

	return versions, nil
}

// MigrateUp applies all pending migrations
func (m *Migrator) MigrateUp() error {
	m.logger.Info("Starting database migration")

	migrations, err := m.LoadMigrations()
	if err != nil {
		return fmt.Errorf("failed to load migrations: %w", err)
	}

	appliedVersions, err := m.GetAppliedMigrations()
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	appliedMap := make(map[int]bool)
	for _, v := range appliedVersions {
		appliedMap[v] = true
	}

	appliedCount := 0

	for _, migration := range migrations {
		if appliedMap[migration.Version] {
			m.logger.Debug("Migration already applied", "version", migration.Version, "description", migration.Description)
			continue
		}

		m.logger.Info("Applying migration", "version", migration.Version, "description", migration.Description)

		// Start transaction
		tx, err := m.db.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}

		// Apply the migration
		if _, err := tx.Exec(migration.UpSQL); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to apply migration %d: %w", migration.Version, err)
		}

		// Record that migration was applied
		if _, err := tx.Exec(
			"INSERT INTO schema_migrations (version, description, applied_at) VALUES (?, ?, ?)",
			migration.Version, migration.Description, time.Now(),
		); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration %d: %w", migration.Version, err)
		}

		// Commit transaction
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit migration %d: %w", migration.Version, err)
		}

		appliedCount++
		m.logger.Info("Migration applied successfully", "version", migration.Version)
	}

	if appliedCount == 0 {
		m.logger.Info("No pending migrations")
	} else {
		m.logger.Info("Database migration completed", "applied_count", appliedCount)
	}

	return nil
}

// MigrateDown rolls back the last N migrations
func (m *Migrator) MigrateDown(steps int) error {
	if steps <= 0 {
		return fmt.Errorf("steps must be greater than 0")
	}

	m.logger.Info("Starting migration rollback", "steps", steps)

	migrations, err := m.LoadMigrations()
	if err != nil {
		return fmt.Errorf("failed to load migrations: %w", err)
	}

	appliedVersions, err := m.GetAppliedMigrations()
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	if len(appliedVersions) == 0 {
		m.logger.Info("No migrations to rollback")
		return nil
	}

	// Sort in descending order for rollback
	sort.Sort(sort.Reverse(sort.IntSlice(appliedVersions)))

	migrationMap := make(map[int]Migration)
	for _, m := range migrations {
		migrationMap[m.Version] = m
	}

	rolledBack := 0
	for _, version := range appliedVersions {
		if rolledBack >= steps {
			break
		}

		migration, exists := migrationMap[version]
		if !exists {
			m.logger.Warn("Migration not found in files", "version", version)
			continue
		}

		if migration.DownSQL == "" {
			return fmt.Errorf("migration %d has no down script", version)
		}

		m.logger.Info("Rolling back migration", "version", version, "description", migration.Description)

		// Start transaction
		tx, err := m.db.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}

		// Apply the rollback
		if _, err := tx.Exec(migration.DownSQL); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to rollback migration %d: %w", version, err)
		}

		// Remove migration record
		if _, err := tx.Exec("DELETE FROM schema_migrations WHERE version = ?", version); err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to remove migration record %d: %w", version, err)
		}

		// Commit transaction
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit rollback %d: %w", version, err)
		}

		rolledBack++
		m.logger.Info("Migration rolled back successfully", "version", version)
	}

	m.logger.Info("Migration rollback completed", "rolled_back_count", rolledBack)
	return nil
}

// Status shows current migration status
func (m *Migrator) Status() error {
	migrations, err := m.LoadMigrations()
	if err != nil {
		return fmt.Errorf("failed to load migrations: %w", err)
	}

	appliedVersions, err := m.GetAppliedMigrations()
	if err != nil {
		return err
	}

	appliedMap := make(map[int]bool)
	for _, v := range appliedVersions {
		appliedMap[v] = true
	}

	fmt.Println("\nMigration Status:")
	fmt.Println("================")

	for _, migration := range migrations {
		status := "PENDING"
		if appliedMap[migration.Version] {
			status = "APPLIED"
		}

		fmt.Printf("Version %03d: %-50s [%s]\n",
			migration.Version,
			migration.Description,
			status)
	}

	fmt.Printf("\nTotal migrations: %d\n", len(migrations))
	fmt.Printf("Applied: %d\n", len(appliedVersions))
	fmt.Printf("Pending: %d\n", len(migrations)-len(appliedVersions))

	return nil
}
