package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

type TaskStatus struct {
	TaskId         int    `json:"task_id"`
	Title          string `json:"title"`
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

func GetTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getting all task statuses in task_history \n")
	query := fmt.Sprintf(`
		SELECT th.id, th.task_id, th.completed, th.completion_date,
		t.id, t.title, t.description, t.status, t.due_date 
		FROM task_history AS th
		JOIN tasks AS t
		ON th.task_id = t.id
		WHERE 1=1
  		AND STRFTIME('%%Y-%%m-%%d', completion_date) = '%s'
	`, time.Now().Format("2006-01-02"))
	fmt.Println(query)
	rows, err := db.GetDB().Query(query)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to get events", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var events []TaskHistory
	var tasks []Task
	var statuses []TaskStatus
	for rows.Next() {
		var event TaskHistory
		var task Task
		if err := rows.Scan(
			&event.ID, &event.TaskId, &event.Completed, &event.CompletionDate,
			&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate,
		); err != nil {
			http.Error(w, "Failed to scan event", http.StatusInternalServerError)
			return
		}
		status := TaskStatus{
			TaskId:         task.ID,
			Title:          task.Title,
			Completed:      event.Completed,
			CompletionDate: event.CompletionDate,
		}
		events = append(events, event)
		tasks = append(tasks, task)
		statuses = append(statuses, status)
	}

	fmt.Printf("got all the statuses\n")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statuses)
}
