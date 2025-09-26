package router

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/ryutaKimu/go_todo/internal/controller"
	"github.com/ryutaKimu/go_todo/internal/infra/postgre/todo"
	service "github.com/ryutaKimu/go_todo/internal/service"
)

func NewRouter(pg *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	repo := todo.NewTodoRepository(pg)
	todoService := service.NewTodoService(repo)
	todoController := controller.NewTodoController(todoService)

	r.Get("/todos", todoController.FetchAllTodoHandler)
	r.Get("/todos/{id}", todoController.FindTodoHandler)
	r.Post("/create", todoController.CreateTodoHandler)
	r.Put("/todos/{id}", todoController.UpdateTodoHandler)
	r.Delete("/todos/{id}", todoController.DeleteHandler)

	return r
}
