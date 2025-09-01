package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Todo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

var todos = []Todo{
	{ID: 1, Title: "todo1", Complete: false},
	{ID: 2, Title: "todo2", Complete: false},
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	var newTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")

	i := findIndex(id)
	if i == -1 {
		http.Error(w, "Method Not Allowed", http.StatusNotFound)
		return
	}

	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = todos[i].ID
	todos[i] = todo
	json.NewEncoder(w).Encode(todo)

}

func findIndex(id string) int {
	tId, err := strconv.Atoi(id)
	if err != nil {
		return -1
	}

	for i, b := range todos {
		if b.ID == tId {
			return i
		}
	}
	return -1
}
