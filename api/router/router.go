package router

import (
	"go_todo_app/handler"

	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/todos", handler.GetTodos)
	r.Post("/todos/add", handler.AddTodo)
	r.Put("/todos/{id}", handler.UpdateTodo)
	return r
}
