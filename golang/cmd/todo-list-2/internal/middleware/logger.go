// ===================================================================
// internal/middleware/logging.go (NEW FILE)
// ===================================================================
package middleware

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"time"
	"todo-list-2/internal/auth"
	"todo-list-2/internal/logger"
)

// Response writer wrapper to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = 200
	}
	return rw.ResponseWriter.Write(b)
}

// LoggingMiddleware logs HTTP requests and responses
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Generate request ID
		requestID := generateRequestID()
		ctx := context.WithValue(r.Context(), logger.RequestIDKey, requestID)
		r = r.WithContext(ctx)

		// Add request ID to response headers for tracing
		w.Header().Set("X-Request-ID", requestID)

		// Wrap response writer to capture status code
		wrapped := &responseWriter{
			ResponseWriter: w,
			statusCode:     0,
		}

		// Get client IP
		clientIP := getClientIP(r)

		// Get user info if available (after auth middleware)
		var userID int
		var username string
		if user, ok := r.Context().Value(UserContextKey).(*auth.Claims); ok {
			userID = user.UserID
			username = user.Username
			ctx = context.WithValue(ctx, logger.UserIDKey, userID)
			ctx = context.WithValue(ctx, logger.UsernameKey, username)
			r = r.WithContext(ctx)
		}

		// Log request
		logger.LogHTTPRequest(r.Method, r.URL.Path, r.UserAgent(), clientIP, userID, username)

		// Process request
		next.ServeHTTP(wrapped, r)

		// Log response
		duration := time.Since(start)
		logger.LogHTTPResponse(r.Method, r.URL.Path, wrapped.statusCode, duration, userID)
	})
}

func generateRequestID() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return fmt.Sprintf("%x", bytes)
}

func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header (for proxies)
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return xff
	}

	// Check X-Real-IP header (for proxies)
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// Fall back to remote address
	return r.RemoteAddr
}
