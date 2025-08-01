package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"my-go-backend/db"
	"my-go-backend/handlers"

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
	db.InitDB(os.Getenv("DATABASE_FILE"))
	defer db.GetDB().Close()

	r := mux.NewRouter()

	r.HandleFunc("/items", handlers.GetItemsHandler).Methods("GET")
	r.HandleFunc("/item", handlers.AddItemHandler).Methods("POST")
	r.HandleFunc("/item/{id:[0-9]+}", handlers.GetItemsHandler).Methods("GET")

	// http.HandleFunc("/add-task", createTaskHandler)
	r.HandleFunc("/get-tasks", handlers.GetTasksHandler)
	// http.HandleFunc("/update-task", updateTaskHandler)
	// http.HandleFunc("/delete-task", deleteTaskHandler)

	r.HandleFunc("/tasks", handlers.GetTasksHandler).Methods("GET")                 // GET /tasks
	r.HandleFunc("/task", handlers.CreateTaskHandler).Methods("POST")               // POST /tasks
	r.HandleFunc("/task/{id:[0-9]+}", handlers.GetTaskHandler).Methods("GET")       // GET /tasks/{id}
	r.HandleFunc("/task/{id:[0-9]+}", handlers.UpdateTaskHandler).Methods("PUT")    // PUT /tasks/{id}
	r.HandleFunc("/task/{id:[0-9]+}", handlers.DeleteTaskHandler).Methods("DELETE") // DELETE /tasks/

	r.HandleFunc("/task-history", handlers.GetTaskHistoryHandler).Methods("GET")   // GET /task-history
	r.HandleFunc("/task-history", handlers.CreateTaskEventHandler).Methods("POST") // POST /task-history
	// http.Handle("/", http.FileServer(http.Dir("./web"))

	r.HandleFunc("/get-tasks-statuses", handlers.GetTaskStatusHandler).Methods("GET") // GET

	// Serve static files from the "web" directory
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/"))))
	http.Handle("/", r)

	port := os.Getenv("PORT")
	fmt.Printf("Server is running on port %v\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
