// ===================================================================
// integration_test.go (NEW FILE) - Integration tests
// ===================================================================
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list-2/internal/auth"
	"todo-list-2/internal/database"
	"todo-list-2/internal/handlers"
	"todo-list-2/internal/middleware"
	"todo-list-2/internal/models"
	"todo-list-2/internal/testutil"

	"github.com/gorilla/mux"
)

func setupTestServer(t *testing.T) (*httptest.Server, *database.DB) {
	// Setup test database
	db := testutil.SetupTestDB(t)

	// Set JWT secret for testing
	auth.SetJWTSecret("test-secret")

	// Setup handlers
	authHandler := handlers.NewAuthHandler(db)
	todoHandler := handlers.NewTodoHandler(db)
	categoryHandler := handlers.NewCategoryHandler(db)

	// Setup routes
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	// Public routes
	api.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	api.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// Protected routes
	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	protected.HandleFunc("/todos", todoHandler.GetAllTodos).Methods("GET")
	protected.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	protected.HandleFunc("/categories", categoryHandler.GetAllCategories).Methods("GET")

	server := httptest.NewServer(r)
	return server, db
}

func TestFullUserFlow(t *testing.T) {
	server, db := setupTestServer(t)
	defer server.Close()
	defer db.Close()

	// Step 1: Register user
	registerReq := models.RegisterRequest{
		Username: "integrationuser",
		Email:    "integration@example.com",
		Password: "password123",
	}

	body, _ := json.Marshal(registerReq)
	resp, err := http.Post(server.URL+"/api/auth/register", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Registration request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status 201, got %d", resp.StatusCode)
	}

	var authResp models.AuthResponse
	json.NewDecoder(resp.Body).Decode(&authResp)
	token := authResp.Token

	// Step 2: Create category
	categoryReq := models.CreateCategoryRequest{
		Name:        "Integration Category",
		Description: "Test category",
		Color:       "#00FF00",
	}

	body, _ = json.Marshal(categoryReq)
	req, _ := http.NewRequest("POST", server.URL+"/api/categories", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Create category request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status 201, got %d", resp.StatusCode)
	}

	// Step 3: Create todo
	todoReq := models.CreateTodoRequest{
		Title:       "Integration Todo",
		Description: "A test todo",
	}

	body, _ = json.Marshal(todoReq)
	req, _ = http.NewRequest("POST", server.URL+"/api/todos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Create todo request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status 201, got %d", resp.StatusCode)
	}

	// Step 4: Get todos
	req, _ = http.NewRequest("GET", server.URL+"/api/todos", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Get todos request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	var todos []models.Todo
	json.NewDecoder(resp.Body).Decode(&todos)

	if len(todos) == 0 {
		t.Error("Expected at least one todo")
	}

	if todos[0].Title != todoReq.Title {
		t.Errorf("Expected todo title %s, got %s", todoReq.Title, todos[0].Title)
	}
}
