// ===================================================================
// internal/config/config.go (UPDATED) - Environment-based DB selection
// ===================================================================
package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	LogLevel    string
	LogFormat   string
	JWTSecret   string
	Environment string
	BasePath    string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getDatabaseURL(),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		LogFormat:   getEnv("LOG_FORMAT", "text"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key-change-this-in-production"),
		Environment: getEnv("ENVIRONMENT", "development"),
		BasePath:    getEnv("BASE_PATH", "."),
	}
}

func getDatabaseURL() string {
	// Check for explicit database URL first TODO
	if url := os.Getenv("DATABASE_URL"); url != "" {
		return url
	}

	// Environment-based defaults
	env := getEnv("ENVIRONMENT", "development")
	switch env {
	case "production":
		return getEnv("DATABASE_URL", "postgres://todouser:todopass@localhost:5432/todos?sslmode=require")
	case "test":
		return getEnv("TEST_DATABASE_URL", ":memory:")
	default: // development
		return getEnv("DATABASE_URL", "data/todos.db")
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
