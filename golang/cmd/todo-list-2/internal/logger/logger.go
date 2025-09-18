// internal/logger/logger.go (NEW FILE)
package logger

import (
	"context"
	"log/slog"
	"os"
	"time"
)

var Logger *slog.Logger

func Init(level string, format string) {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level:     logLevel,
		AddSource: true,
	}

	switch format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	Logger = slog.New(handler)
	slog.SetDefault(Logger)
}

// Context keys for request logging
type contextKey string

const (
	RequestIDKey contextKey = "request_id"
	UserIDKey    contextKey = "user_id"
	UsernameKey  contextKey = "username"
)

// WithRequestContext adds request context to logger
func WithRequestContext(ctx context.Context) *slog.Logger {
	logger := Logger

	if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
		logger = logger.With("request_id", requestID)
	}

	if userID, ok := ctx.Value(UserIDKey).(int); ok {
		logger = logger.With("user_id", userID)
	}

	if username, ok := ctx.Value(UsernameKey).(string); ok {
		logger = logger.With("username", username)
	}

	return logger
}

// WithUser adds user context to logger
func WithUser(userID int, username string) *slog.Logger {
	return Logger.With(
		"user_id", userID,
		"username", username,
	)
}

// HTTP request logging helpers
func LogHTTPRequest(method, path, userAgent, clientIP string, userID int, username string) {
	Logger.Info("HTTP request",
		"method", method,
		"path", path,
		"user_agent", userAgent,
		"client_ip", clientIP,
		"user_id", userID,
		"username", username,
	)
}

func LogHTTPResponse(method, path string, statusCode int, duration time.Duration, userID int) {
	level := slog.LevelInfo
	if statusCode >= 400 && statusCode < 500 {
		level = slog.LevelWarn
	} else if statusCode >= 500 {
		level = slog.LevelError
	}

	Logger.Log(context.Background(), level, "HTTP response",
		"method", method,
		"path", path,
		"status_code", statusCode,
		"duration_ms", duration.Milliseconds(),
		"user_id", userID,
	)
}

// Database operation logging
func LogDBOperation(operation, table string, userID int, duration time.Duration, err error) {
	if err != nil {
		Logger.Error("Database operation failed",
			"operation", operation,
			"table", table,
			"user_id", userID,
			"duration_ms", duration.Milliseconds(),
			"error", err.Error(),
		)
	} else {
		Logger.Debug("Database operation completed",
			"operation", operation,
			"table", table,
			"user_id", userID,
			"duration_ms", duration.Milliseconds(),
		)
	}
}

// Authentication logging
func LogAuthSuccess(username string, userID int, action string) {
	Logger.Info("Authentication successful",
		"action", action,
		"username", username,
		"user_id", userID,
	)
}

func LogAuthFailure(username string, action string, reason string, clientIP string) {
	Logger.Warn("Authentication failed",
		"action", action,
		"username", username,
		"reason", reason,
		"client_ip", clientIP,
	)
}

// Application lifecycle logging
func LogStartup(port string) {
	Logger.Info("Server starting",
		"port", port,
		"version", "1.0.0", // You can make this configurable
	)
}

func LogShutdown() {
	Logger.Info("Server shutting down")
}
