// // ===================================================================
// // tests/database_test.go (NEW FILE) - Database tests in separate package
// // ===================================================================
package tests

// import (
// 	"testing"
// 	"todo-list-2/internal/database"
// 	"todo-list-2/internal/models"
// 	"todo-list-2/internal/testutil"
// )

// func TestCreateUser(t *testing.T) {
// 	// Initialize logger
// 	logger := logger.Init("info", "text")
// 	logger.LogStartup("8080")

// 	testDB := testutil.SetupTestDB(t)
// 	defer testDB.Close()

// 	// Create the database.DB wrapper
// 	db := &database.DB{conn: testDB.Conn, logger: logger}

// 	tests := []struct {
// 		name    string
// 		req     models.RegisterRequest
// 		wantErr bool
// 	}{
// 		{
// 			name: "valid user",
// 			req: models.RegisterRequest{
// 				Username: "john",
// 				Email:    "john@example.com",
// 				Password: "password123",
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "duplicate username",
// 			req: models.RegisterRequest{
// 				Username: "john", // Same as above
// 				Email:    "john2@example.com",
// 				Password: "password123",
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			user, err := db.CreateUser(tt.req)

// 			if tt.wantErr {
// 				if err == nil {
// 					t.Error("Expected error but got none")
// 				}
// 				return
// 			}

// 			if err != nil {
// 				t.Errorf("Unexpected error: %v", err)
// 				return
// 			}

// 			if user.Username != tt.req.Username {
// 				t.Errorf("Expected username %s, got %s", tt.req.Username, user.Username)
// 			}

// 			if user.Email != tt.req.Email {
// 				t.Errorf("Expected email %s, got %s", tt.req.Email, user.Email)
// 			}

// 			if user.ID == 0 {
// 				t.Error("Expected user ID to be set")
// 			}
// 		})
// 	}
// }

// func TestAuthenticateUser(t *testing.T) {
// 	testDB := testutil.SetupTestDB(t)
// 	defer testDB.Close()

// 	db := &database.DB{testDB.DB}

// 	// Create a test user
// 	testUser := testutil.CreateTestUser(t, testDB)

// 	tests := []struct {
// 		name     string
// 		username string
// 		password string
// 		wantErr  bool
// 	}{
// 		{
// 			name:     "valid credentials",
// 			username: testUser.Username,
// 			password: "password123",
// 			wantErr:  false,
// 		},
// 		{
// 			name:     "invalid password",
// 			username: testUser.Username,
// 			password: "wrongpassword",
// 			wantErr:  true,
// 		},
// 		{
// 			name:     "invalid username",
// 			username: "nonexistent",
// 			password: "password123",
// 			wantErr:  true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			user, err := db.AuthenticateUser(tt.username, tt.password)

// 			if tt.wantErr {
// 				if err == nil {
// 					t.Error("Expected error but got none")
// 				}
// 				return
// 			}

// 			if err != nil {
// 				t.Errorf("Unexpected error: %v", err)
// 				return
// 			}

// 			if user.Username != tt.username {
// 				t.Errorf("Expected username %s, got %s", tt.username, user.Username)
// 			}
// 		})
// 	}
// }

// func TestTodoOperations(t *testing.T) {
// 	testDB := testutil.SetupTestDB(t)
// 	defer testDB.Close()

// 	db := &database.DB{testDB.DB}

// 	// Setup test data
// 	user := testutil.CreateTestUser(t, testDB)
// 	category := testutil.CreateTestCategory(t, testDB, user.ID)

// 	t.Run("create todo", func(t *testing.T) {
// 		req := models.CreateTodoRequest{
// 			Title:       "Test Todo",
// 			Description: "Test Description",
// 			CategoryID:  &category.ID,
// 		}

// 		todo, err := db.CreateTodo(user.ID, req)
// 		if err != nil {
// 			t.Fatalf("Failed to create todo: %v", err)
// 		}

// 		if todo.Title != req.Title {
// 			t.Errorf("Expected title %s, got %s", req.Title, todo.Title)
// 		}

// 		if todo.UserID != user.ID {
// 			t.Errorf("Expected user_id %d, got %d", user.ID, todo.UserID)
// 		}

// 		if todo.Category == nil || todo.Category.ID != category.ID {
// 			t.Error("Expected category to be populated")
// 		}
// 	})

// 	t.Run("get user todos", func(t *testing.T) {
// 		// Create multiple todos
// 		testutil.CreateTestTodo(t, testDB, user.ID, &category.ID)
// 		testutil.CreateTestTodo(t, testDB, user.ID, nil) // No category

// 		todos, err := db.GetAllTodos(user.ID)
// 		if err != nil {
// 			t.Fatalf("Failed to get todos: %v", err)
// 		}

// 		if len(todos) < 2 {
// 			t.Errorf("Expected at least 2 todos, got %d", len(todos))
// 		}

// 		// All todos should belong to the user
// 		for _, todo := range todos {
// 			if todo.UserID != user.ID {
// 				t.Errorf("Todo belongs to wrong user: expected %d, got %d", user.ID, todo.UserID)
// 			}
// 		}
// 	})

// 	t.Run("update todo", func(t *testing.T) {
// 		todo := testutil.CreateTestTodo(t, testDB, user.ID, nil)

// 		newTitle := "Updated Title"
// 		completed := true

// 		req := models.UpdateTodoRequest{
// 			Title:     &newTitle,
// 			Completed: &completed,
// 		}

// 		updated, err := db.UpdateTodo(user.ID, todo.ID, req)
// 		if err != nil {
// 			t.Fatalf("Failed to update todo: %v", err)
// 		}

// 		if updated.Title != newTitle {
// 			t.Errorf("Expected title %s, got %s", newTitle, updated.Title)
// 		}

// 		if !updated.Completed {
// 			t.Error("Expected todo to be completed")
// 		}
// 	})

// 	t.Run("delete todo", func(t *testing.T) {
// 		todo := testutil.CreateTestTodo(t, testDB, user.ID, nil)

// 		err := db.DeleteTodo(user.ID, todo.ID)
// 		if err != nil {
// 			t.Fatalf("Failed to delete todo: %v", err)
// 		}

// 		// Try to get the deleted todo
// 		_, err = db.GetTodoByID(user.ID, todo.ID)
// 		if err == nil {
// 			t.Error("Expected error when getting deleted todo")
// 		}
// 	})
// }

// func TestUserIsolation(t *testing.T) {
// 	testDB := testutil.SetupTestDB(t)
// 	defer testDB.Close()

// 	db := &database.DB{testDB.DB}

// 	// Create two users
// 	user1 := testutil.CreateTestUser(t, testDB)

// 	user2Req := models.RegisterRequest{
// 		Username: "user2",
// 		Email:    "user2@example.com",
// 		Password: "password123",
// 	}
// 	user2, err := db.CreateUser(user2Req)
// 	if err != nil {
// 		t.Fatalf("Failed to create user2: %v", err)
// 	}

// 	// Create todos for each user
// 	testutil.CreateTestTodo(t, testDB, user1.ID, nil)
// 	testutil.CreateTestTodo(t, testDB, user2.ID, nil)

// 	// User1 should only see their todos
// 	user1Todos, err := db.GetAllTodos(user1.ID)
// 	if err != nil {
// 		t.Fatalf("Failed to get user1 todos: %v", err)
// 	}

// 	for _, todo := range user1Todos {
// 		if todo.UserID != user1.ID {
// 			t.Error("User1 can see other user's todos")
// 		}
// 	}

// 	// User2 should only see their todos
// 	user2Todos, err := db.GetAllTodos(user2.ID)
// 	if err != nil {
// 		t.Fatalf("Failed to get user2 todos: %v", err)
// 	}

// 	for _, todo := range user2Todos {
// 		if todo.UserID != user2.ID {
// 			t.Error("User2 can see other user's todos")
// 		}
// 	}
// }
