// internal/logger/logger.go (NEW FILE)
package logger

import (
	"context"
	"log/slog"
	"os"
	"time"
)

type Logger struct {
	*slog.Logger
}

func Init(level string, format string) *Logger {
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

	logger := slog.New(handler)
	slog.SetDefault(logger)
	return &Logger{logger}
}

// // Convenience wrappers
// func Debug(msg string, args ...any) {
// 	Logger.Debug(msg, args...)
// }

// func Info(msg string, args ...any) {
// 	Logger.Info(msg, args...)
// }

// func Warn(msg string, args ...any) {
// 	Logger.Warn(msg, args...)
// }

// func Error(msg string, args ...any) {
// 	Logger.Error(msg, args...)
// }

// Context keys for request logging
type contextKey string

const (
	RequestIDKey contextKey = "request_id"
	UserIDKey    contextKey = "user_id"
	UsernameKey  contextKey = "username"
)

// WithRequestContext adds request context to logger
func (l *Logger) WithRequestContext(ctx context.Context) *Logger {
	var tempLogger *slog.Logger
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
		tempLogger = l.With("request_id", requestID)
	}

	if userID, ok := ctx.Value(UserIDKey).(int); ok {
		tempLogger = l.With("user_id", userID)
	}

	if username, ok := ctx.Value(UsernameKey).(string); ok {
		tempLogger = l.With("username", username)
	}

	return &Logger{tempLogger}
}

// WithUser adds user context to logger
func (l *Logger) WithUser(userID int, username string) *Logger {
	return &Logger{l.With(
		"user_id", userID,
		"username", username,
	)}
}

// HTTP request logging helpers
func (l *Logger) LogHTTPRequest(method, path, userAgent, clientIP string, userID int, username string) {
	l.Info("HTTP request",
		"method", method,
		"path", path,
		"user_agent", userAgent,
		"client_ip", clientIP,
		"user_id", userID,
		"username", username,
	)
}

func (l *Logger) LogHTTPResponse(method, path string, statusCode int, duration time.Duration, userID int) {
	level := slog.LevelInfo
	if statusCode >= 400 && statusCode < 500 {
		level = slog.LevelWarn
	} else if statusCode >= 500 {
		level = slog.LevelError
	}

	l.Log(context.Background(), level, "HTTP response",
		"method", method,
		"path", path,
		"status_code", statusCode,
		"duration_ms", duration.Milliseconds(),
		"user_id", userID,
	)
}

// Database operation logging
func (l *Logger) LogDBOperation(operation, table string, userID int, duration time.Duration, err error) {
	if err != nil {
		l.Error("Database operation failed",
			"operation", operation,
			"table", table,
			"user_id", userID,
			"duration_ms", duration.Milliseconds(),
			"error", err.Error(),
		)
	} else {
		l.Debug("Database operation completed",
			"operation", operation,
			"table", table,
			"user_id", userID,
			"duration_ms", duration.Milliseconds(),
		)
	}
}

// Authentication logging
func (l *Logger) LogAuthSuccess(username string, userID int, action string) {
	l.Info("Authentication successful",
		"action", action,
		"username", username,
		"user_id", userID,
	)
}

func (l *Logger) LogAuthFailure(username string, action string, reason string, clientIP string) {
	l.Warn("Authentication failed",
		"action", action,
		"username", username,
		"reason", reason,
		"client_ip", clientIP,
	)
}

// Application lifecycle logging
func (l *Logger) LogStartup(port string) {
	l.Info("Server starting",
		"port", port,
		"version", "1.0.0", // You can make this configurable
	)
}

func (l *Logger) LogShutdown() {
	l.Info("Server shutting down")
}
