// ===================================================================
// cmd/server/main.go (UPDATED with logging and config)
// ===================================================================
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-list-2/internal/auth"
	"todo-list-2/internal/config"
	"todo-list-2/internal/database"
	"todo-list-2/internal/handlers"
	"todo-list-2/internal/logger"
	"todo-list-2/internal/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	logger.Init(cfg.LogLevel, cfg.LogFormat)
	logger.LogStartup(cfg.Port)

	// Set JWT secret
	auth.SetJWTSecret(cfg.JWTSecret)

	// Initialize database
	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		logger.Logger.Error("Failed to connect to database", "error", err.Error())
		os.Exit(1)
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Logger.Error("Failed to close database", "error", err.Error())
		}
	}()

	logger.Logger.Info("Database connected successfully")

	// Initialize handlers
	todoHandler := handlers.NewTodoHandler(db)
	categoryHandler := handlers.NewCategoryHandler(db)
	authHandler := handlers.NewAuthHandler(db)

	// Setup routes
	r := mux.NewRouter()

	// Add logging middleware to all routes
	r.Use(middleware.LoggingMiddleware)

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Public auth routes (no authentication required)
	api.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	api.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// Health check endpoint
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","timestamp":"` + time.Now().Format(time.RFC3339) + `"}`))
	}).Methods("GET")

	// Protected routes (require authentication)
	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// User info route
	protected.HandleFunc("/auth/me", authHandler.Me).Methods("GET")

	// Todo routes (all protected)
	protected.HandleFunc("/todos", todoHandler.GetAllTodos).Methods("GET")
	protected.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	protected.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	protected.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")

	// Category routes (all protected)
	protected.HandleFunc("/categories", categoryHandler.GetAllCategories).Methods("GET")
	protected.HandleFunc("/categories", categoryHandler.CreateCategory).Methods("POST")
	protected.HandleFunc("/categories/{id}", categoryHandler.UpdateCategory).Methods("PUT")
	protected.HandleFunc("/categories/{id}", categoryHandler.DeleteCategory).Methods("DELETE")

	// Serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(r)

	// Setup server with graceful shutdown
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		logger.Logger.Info("Server listening", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Error("Failed to start server", "error", err.Error())
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.LogShutdown()

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Error("Server forced to shutdown", "error", err.Error())
		os.Exit(1)
	}

	logger.Logger.Info("Server exited successfully")
}
