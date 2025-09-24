package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ryutaKimu/go_todo/internal/controller/services"
	"github.com/ryutaKimu/go_todo/internal/model"
)

type TodoController struct {
	service services.TodoService
}

func NewTodoController(service services.TodoService) *TodoController {
	return &TodoController{service: service}
}

func (c *TodoController) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.service.CreateTodo(r.Context(), &todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
