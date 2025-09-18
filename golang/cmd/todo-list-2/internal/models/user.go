// First, install additional dependencies:
// go get golang.org/x/crypto/bcrypt
// go get github.com/golang-jwt/jwt/v5

// ===================================================================
// internal/models/user.go (NEW FILE)
// ===================================================================
package models

import "time"

type User struct {
    ID        int       `json:"id" db:"id"`
    Username  string    `json:"username" db:"username"`
    Email     string    `json:"email" db:"email"`
    Password  string    `json:"-" db:"password"` // Never return password in JSON
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type RegisterRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type AuthResponse struct {
    User  User   `json:"user"`
    Token string `json:"token"`
}
