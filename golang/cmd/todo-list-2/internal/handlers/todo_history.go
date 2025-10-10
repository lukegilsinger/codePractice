package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-list-2/internal/database"
	"todo-list-2/internal/middleware"
	"todo-list-2/internal/models"

	"github.com/gorilla/mux"
)

type TodoHistoryHandler struct {
	db *database.DB
}

func NewTodoHistoryHandler(db *database.DB) *TodoHistoryHandler {
	return &TodoHistoryHandler{db: db}
}

// CompleteTodo marks a task as complete and records it in history
func (h *TodoHistoryHandler) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUserFromContext(r)
	if user == nil {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	var req models.CreateTodoCompletionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Create completion record
	completion, err := h.db.CreateTodoCompletion(user.UserID, todoID, req.CompletedAt, req.Notes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// // Update the todo (for recurring tasks, calculate next due date)
	// todo, err := h.db.GetTodoByID(user.UserID, todoID)
	// if err != nil {
	// 	http.Error(w, "Todo not found", http.StatusNotFound)
	// 	return
	// }

	// // Update todo based on frequency
	// updateReq := models.UpdateTodoRequest{}
	// if todo.Frequency == models.FrequencyOnce {
	// 	completed := true
	// 	updateReq.Completed = &completed
	// } else {
	// 	// Recurring task - calculate next due date
	// 	nextDue := models.CalculateNextDueDate(todo.Frequency, time.Now())
	// 	updateReq.NextDueDate = nextDue
	// }

	// updatedTodo, err := h.db.UpdateTodo(user.UserID, todoID, updateReq)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// response := map[string]interface{}{
	// 	"completion": completion,
	// 	"todo":       updatedTodo,
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completion)
}

// GetTaskHistory gets completion history for a task
func (h *TodoHistoryHandler) GetTodoHistory(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUserFromContext(r)
	if user == nil {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	history, err := h.db.GetTodoHistory(user.UserID, todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// // Get statistics
	// count, _ := h.db.GetTaskCompletionCount(user.UserID, todoID)
	// streak, _ := h.db.GetCurrentStreak(user.UserID, todoID)
	// completedToday, _ := h.db.IsTaskCompletedToday(user.UserID, todoID)

	// response := map[string]interface{}{
	// 	"completions":     history,
	// 	"total_count":     count,
	// 	"current_streak":  streak,
	// 	"completed_today": completedToday,
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}
