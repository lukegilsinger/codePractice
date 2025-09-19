// ===================================================================
// internal/auth/jwt_test.go (NEW FILE) - JWT tests
// ===================================================================
package auth

import (
	"testing"
	"time"
)

func TestJWTOperations(t *testing.T) {
	// Set test secret
	SetJWTSecret("test-secret-key")

	userID := 123
	username := "testuser"

	t.Run("generate and validate token", func(t *testing.T) {
		token, err := GenerateToken(userID, username)
		if err != nil {
			t.Fatalf("Failed to generate token: %v", err)
		}

		if token == "" {
			t.Error("Expected token to be non-empty")
		}

		claims, err := ValidateToken(token)
		if err != nil {
			t.Fatalf("Failed to validate token: %v", err)
		}

		if claims.UserID != userID {
			t.Errorf("Expected user ID %d, got %d", userID, claims.UserID)
		}

		if claims.Username != username {
			t.Errorf("Expected username %s, got %s", username, claims.Username)
		}
	})

	t.Run("invalid token", func(t *testing.T) {
		invalidToken := "invalid.token.here"

		_, err := ValidateToken(invalidToken)
		if err == nil {
			t.Error("Expected error for invalid token")
		}
	})

	t.Run("empty token", func(t *testing.T) {
		_, err := ValidateToken("")
		if err == nil {
			t.Error("Expected error for empty token")
		}
	})
}

func TestTokenExpiration(t *testing.T) {
	SetJWTSecret("test-secret-key")

	// This test would need to mock time or use a shorter expiration
	// For now, we'll just test that the token contains expiration info
	token, err := GenerateToken(123, "testuser")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
		t.Error("Expected token to have valid expiration time")
	}
}
