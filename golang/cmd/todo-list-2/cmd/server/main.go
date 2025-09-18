// cmd/server/main.go (UPDATED with authentication)
package main

import (
	"log"
	"net/http"
	"todo-list-2/internal/database"
	"todo-list-2/internal/handlers"
	"todo-list-2/internal/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize database
	db, err := database.New("todos.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize handlers
	todoHandler := handlers.NewTodoHandler(db)
	categoryHandler := handlers.NewCategoryHandler(db)
	authHandler := handlers.NewAuthHandler(db) // NEW

	// Setup routes
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Public auth routes (no authentication required)
	api.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	api.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

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

	log.Println("Server starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", handler))
}
