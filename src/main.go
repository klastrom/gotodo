package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var (
	todos  = []Todo{}
	nextID = 1
	mu     sync.Mutex
)

func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.ID = nextID
	nextID++
	todos = append(todos, todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, t := range todos {
		if t.ID == todo.ID {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/todos", getTodosHandler)
	http.HandleFunc("/todos/add", addTodoHandler)
	http.HandleFunc("/todos/delete", deleteTodoHandler)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}