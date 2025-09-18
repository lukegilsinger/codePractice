package models

import "time"

type Todo struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Completed   bool      `json:"completed" db:"completed"`
	CategoryID  *int      `json:"category_id" db:"category_id"` // NEW: nullable foreign key
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`

	// Include category info when fetching todos
	Category *Category `json:"category,omitempty"` // NEW: populated in queries
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryID  *int   `json:"category_id,omitempty"` // NEW: optional category
}

type UpdateTodoRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Completed   *bool   `json:"completed,omitempty"`
	CategoryID  *int    `json:"category_id,omitempty"` // NEW: can change category
}
