package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"my-go-backend/db"
	// "my-go-backend/service"

	_ "github.com/mattn/go-sqlite3"
	// "google.golang.org/protobuf/proto"
	// pb "my-go-backend/pb"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("adding item\n")
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
	_, err := db.GetDB().Exec("INSERT INTO items (name) VALUES (?)", item.Name)
	if err != nil {
		http.Error(w, "Failed to add item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getting all items\n")
	rows, err := db.GetDB().Query("SELECT id, name FROM items")
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

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getting an item\n")
	rows, err := db.GetDB().Query("SELECT id, name FROM items")
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
