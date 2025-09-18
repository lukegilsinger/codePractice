// ===================================================================
// internal/config/config.go (NEW FILE)
// ===================================================================
package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port        string
	DatabaseURL string
	LogLevel    string
	LogFormat   string
	JWTSecret   string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "todos.db"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		LogFormat:   getEnv("LOG_FORMAT", "text"), // "text" or "json"
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key-change-this-in-production"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
