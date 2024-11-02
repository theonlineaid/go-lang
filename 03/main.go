package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Item represents a simple item with an ID and a name
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	items  = []Item{}
	nextID = 1
)

func main() {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItemByID).Methods("GET")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}

// Handlers

// Create a new item
func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item.ID = nextID
	nextID++
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

// Get all items
func getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

// Get a single item by ID
func getItemByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, item := range items {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

// Update an item by ID
func updateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, item := range items {
		if item.ID == id {
			var updatedItem Item
			if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			updatedItem.ID = id
			items[i] = updatedItem
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

// Delete an item by ID
func deleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}
