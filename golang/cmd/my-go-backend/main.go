package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	// "github.com/docker/docker/daemon/logger"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	// "google.golang.org/protobuf/proto"
	// pb "my-go-backend/pb"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date"`
}

type Config struct {
	Env string `env:"ENV,required"`
}

var db *sql.DB

func main() {
	// Load environment variables from .env file
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	// Determine the environment and load the corresponding .env file
	env := os.Getenv("GO_ENV") // Set this environment variable to "development", "staging", or "production"
	if env == "" {
		env = "dev" // Default to development if not set
	}

	envFile := "environments/" + env + ".env"
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Printf("No %s file found, using default environment variables", envFile)
	}

	fmt.Println("ENV: " + os.Getenv("ENV"))
	// cfg := Config{} // 👈 new instance of `Config`

	// err = env.Parse(&cfg) // 👈 Parse environment variables into `Config`
	// if err != nil {
	// 	log.Fatalf("unable to parse ennvironment variables: %e", err)
	// }

	logger.Info("Info message", "key2", "value2")
	initDB(os.Getenv("DATABASE_FILE"))
	defer db.Close()

	r := mux.NewRouter()

	// http.HandleFunc("/add-item", addItemHandler)
	// http.HandleFunc("/get-items", getItemsHandler)

	// http.HandleFunc("/add-task", createTaskHandler)
	http.HandleFunc("/get-tasks", getTasksHandler)
	// http.HandleFunc("/update-task", updateTaskHandler)
	// http.HandleFunc("/delete-task", deleteTaskHandler)

	r.HandleFunc("/tasks", getTasksHandler).Methods("GET")                 // GET /tasks
	r.HandleFunc("/task", createTaskHandler).Methods("POST")               // POST /tasks
	r.HandleFunc("/task/{id:[0-9]+}", getTaskHandler).Methods("GET")       // GET /tasks/{id}
	r.HandleFunc("/task/{id:[0-9]+}", updateTaskHandler).Methods("PUT")    // PUT /tasks/{id}
	r.HandleFunc("/task/{id:[0-9]+}", deleteTaskHandler).Methods("DELETE") // DELETE /tasks/

	// http.Handle("/", http.FileServer(http.Dir("./web"))

	// Serve static files from the "web" directory
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/"))))
	http.Handle("/", r)

	port := os.Getenv("PORT")
	fmt.Printf("Server is running on port %v\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func initDB(databaseFile string) {
	fmt.Println("initializing sqlite db from " + databaseFile)
	var err error

	db, err = sql.Open("sqlite3", "./"+databaseFile)
	if err != nil {
		panic(err)
	}

	// Create tasks table if it doesn't exist
	sqlStmt := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT,
        status TEXT NOT NULL,
        due_date DATE
    );`
	if _, err = db.Exec(sqlStmt); err != nil {
		panic(err)
	}
	fmt.Println("sqlite db is running")
}

func addItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Insert the item into the database
	_, err := db.Exec("INSERT INTO items (name) VALUES (?)", item.Name)
	if err != nil {
		http.Error(w, "Failed to add item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getItemsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM items")
	if err != nil {
		http.Error(w, "Failed to get items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			http.Error(w, "Failed to scan item", http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
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
	result, err := db.Exec("INSERT INTO tasks (title, description, status, due_date) VALUES (?, ?, ?, ?)",
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

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, description, status, due_date FROM tasks")
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

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getting a task\n")
	idStr := r.URL.Path[len("/task/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT id, title, description, status, due_date FROM tasks WHERE id = ?", id)
	var task Task
	err = row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate)
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

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
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
	_, err = db.Exec("UPDATE tasks SET title = ?, description = ?, status = ?, due_date = ? WHERE id = ?",
		task.Title, task.Description, task.Status, task.DueDate, id)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	fmt.Printf("udpated task %v\n", id)

	w.WriteHeader(http.StatusNoContent)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
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
	_, err = db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	fmt.Printf("deleted task %v\n", id)

	w.WriteHeader(http.StatusNoContent)
}
