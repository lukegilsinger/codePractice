package main

import (
	"log"
	"net/http"
	"todo-list-2/internal/database"
	"todo-list-2/internal/handlers"

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
	categoryHandler := handlers.NewCategoryHandler(db) // NEW

	// Setup routes
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Todo routes
	api.HandleFunc("/todos", todoHandler.GetAllTodos).Methods("GET")
	api.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	api.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	api.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")

	// Category routes (NEW)
	api.HandleFunc("/categories", categoryHandler.GetAllCategories).Methods("GET")
	api.HandleFunc("/categories", categoryHandler.CreateCategory).Methods("POST")
	api.HandleFunc("/categories/{id}", categoryHandler.UpdateCategory).Methods("PUT")
	api.HandleFunc("/categories/{id}", categoryHandler.DeleteCategory).Methods("DELETE")

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
