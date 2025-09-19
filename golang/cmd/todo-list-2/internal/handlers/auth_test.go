// ===================================================================
// internal/handlers/auth_test.go (NEW FILE) - Handler tests
// ===================================================================
package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list-2/internal/models"
	"todo-list-2/internal/testutil"
)

func TestAuthHandler_Register(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer db.Close()

	handler := NewAuthHandler(db)

	tests := []struct {
		name           string
		req            models.RegisterRequest
		expectedStatus int
	}{
		{
			name: "valid registration",
			req: models.RegisterRequest{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "invalid email",
			req: models.RegisterRequest{
				Username: "testuser2",
				Email:    "", // Empty email
				Password: "password123",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "short password",
			req: models.RegisterRequest{
				Username: "testuser3",
				Email:    "test3@example.com",
				Password: "12345", // Too short
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.req)
			req := httptest.NewRequest("POST", "/auth/register", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler.Register(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedStatus == http.StatusCreated {
				var response models.AuthResponse
				err := json.NewDecoder(rr.Body).Decode(&response)
				if err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}

				if response.User.Username != tt.req.Username {
					t.Errorf("Expected username %s, got %s", tt.req.Username, response.User.Username)
				}

				if response.Token == "" {
					t.Error("Expected token to be set")
				}
			}
		})
	}
}

func TestAuthHandler_Login(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer db.Close()

	// Create a test user
	testUser := testutil.CreateTestUser(t, db)

	handler := NewAuthHandler(db)

	tests := []struct {
		name           string
		req            models.LoginRequest
		expectedStatus int
	}{
		{
			name: "valid login",
			req: models.LoginRequest{
				Username: testUser.Username,
				Password: "password123",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "invalid password",
			req: models.LoginRequest{
				Username: testUser.Username,
				Password: "wrongpassword",
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "nonexistent user",
			req: models.LoginRequest{
				Username: "nonexistent",
				Password: "password123",
			},
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.req)
			req := httptest.NewRequest("POST", "/auth/login", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler.Login(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var response models.AuthResponse
				err := json.NewDecoder(rr.Body).Decode(&response)
				if err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}

				if response.Token == "" {
					t.Error("Expected token to be set")
				}
			}
		})
	}
}
