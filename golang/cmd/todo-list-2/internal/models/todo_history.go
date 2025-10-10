package models

import "time"

type TodoHistory struct {
	ID          int       `json:"id" db:"id"`
	TodoID      int       `json:"todo_id" db:"todo_id"`
	UserID      int       `json:"user_id" db:"user_id"`
	CompletedAt time.Time `json:"completed_at" db:"completed_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Notes       string    `json:"notes" db:"notes"`
}

type CreateTodoCompletionRequest struct {
	CompletedAt time.Time `json:"completed_at" db:"completed_at"`
	Notes       string    `json:"notes" db:"notes"`
}

// TodoWithHistory includes completion history
type TodoWithHistory struct {
	Todo
	CompletionCount   int           `json:"completion_count"`
	RecentCompletions []TodoHistory `json:"recent_completions"`
	CompletedToday    bool          `json:"completed_today"`
	CurrentStreak     int           `json:"current_streak"`
	LongestStreak     int           `json:"longest_streak"`
}
