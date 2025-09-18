// ===================================================================
// internal/models/todo.go (UPDATED - add UserID)
// ===================================================================
package models

import "time"

type Todo struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"` // NEW: foreign key to users
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Completed   bool      `json:"completed" db:"completed"`
	CategoryID  *int      `json:"category_id" db:"category_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`

	Category *Category `json:"category,omitempty"`
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryID  *int   `json:"category_id,omitempty"`
	// UserID will come from JWT token, not request body
}

type UpdateTodoRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Completed   *bool   `json:"completed,omitempty"`
	CategoryID  *int    `json:"category_id,omitempty"`
}
