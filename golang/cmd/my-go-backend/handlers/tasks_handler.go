package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"my-go-backend/db"
	"my-go-backend/service"

	_ "github.com/mattn/go-sqlite3"
	// "google.golang.org/protobuf/proto"
	// pb "my-go-backend/pb"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Insert the task into the database
	result, err := db.GetDB().Exec("INSERT INTO tasks (title, description, status, due_date) VALUES (?, ?, ?, ?)",
		task.Title, task.Description, task.Status, task.DueDate)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve task ID", http.StatusInternalServerError)
		return
	}

	task.ID = int(id)

	fmt.Printf("created task %v\n", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getting all tasks\n")
	rows, err := db.GetDB().Query("SELECT id, title, description, status, due_date FROM tasks")
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// var taskList pb.TaskList
	var tasks []Task
	for rows.Next() {
		// var task pb.Task
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate); err != nil {
			http.Error(w, "Failed to scan task", http.StatusInternalServerError)
			return
		}
		// taskList.Tasks = append(taskList.Tasks, &task)
		tasks = append(tasks, task)
	}

	fmt.Printf("got all the tasks\n")

	// Serialize the response to protobuf
	// responseData, err := proto.Marshal(&taskList)
	// if err != nil {
	// 	http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/octet-stream")
	// w.WriteHeader(http.StatusOK)
	// w.Write(responseData)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getting a task\n")

	idStr := r.URL.Path[len("/task/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	params := service.GetParams{
		Id:      id,
		Table:   "tasks",
		Columns: []string{"id", "title", "description", "status", "due_date"},
	}

	res, err := service.Get(params)

	var task Task
	err = res.Item.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get task", http.StatusInternalServerError)
		return
	}

	fmt.Printf("got task %v\n", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/tasku/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Update the task in the database
	_, err = db.GetDB().Exec("UPDATE tasks SET title = ?, description = ?, status = ?, due_date = ? WHERE id = ?",
		task.Title, task.Description, task.Status, task.DueDate, id)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	fmt.Printf("udpated task %v\n", id)

	w.WriteHeader(http.StatusNoContent)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/taskd/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// Delete the task from the database
	_, err = db.GetDB().Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	fmt.Printf("deleted task %v\n", id)

	w.WriteHeader(http.StatusNoContent)
}
