package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ryutaKimu/go_todo/internal/controller/services"
	"github.com/ryutaKimu/go_todo/internal/model"
)

type TodoController struct {
	service services.TodoService
}

func NewTodoController(service services.TodoService) *TodoController {
	return &TodoController{service: service}
}

func (c *TodoController) FetchAllTodoHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := c.service.FetchAllTodo(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *TodoController) FindTodoHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	todo, err := c.service.FindTodoById(r.Context(), userId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
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

func (c *TodoController) UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	var todo model.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.service.UpdateTodo(r.Context(), userId, &todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (c *TodoController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	if err := c.service.DeleteTodo(r.Context(), userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
