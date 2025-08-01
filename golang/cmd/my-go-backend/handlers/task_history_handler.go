package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"my-go-backend/db"

	_ "github.com/mattn/go-sqlite3"
	// "google.golang.org/protobuf/proto"
	// pb "my-go-backend/pb"
)

type TaskHistory struct {
	ID             int    `json:"id"`
	TaskId         int    `json:"task_id"`
	Completed      bool   `json:"completed"`
	CompletionDate string `json:"completion_date"`
}

func CreateTaskEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var event TaskHistory
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Insert the task into the database
	result, err := db.GetDB().Exec("INSERT INTO task_history (task_id, completed, completion_date) VALUES (?, ?, ?)",
		event.TaskId, event.Completed, event.CompletionDate)
	if err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve event ID", http.StatusInternalServerError)
		return
	}

	event.ID = int(id)

	fmt.Printf("created event %v in task_history\n", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func GetTaskHistoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getting all events in task_history \n")
	rows, err := db.GetDB().Query("SELECT id, task_id, completed, completion_date FROM task_history")
	if err != nil {
		http.Error(w, "Failed to get events", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var events []TaskHistory
	for rows.Next() {
		var event TaskHistory
		if err := rows.Scan(&event.ID, &event.TaskId, &event.Completed, &event.CompletionDate); err != nil {
			http.Error(w, "Failed to scan event", http.StatusInternalServerError)
			return
		}
		events = append(events, event)
	}

	fmt.Printf("got all the events\n")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
