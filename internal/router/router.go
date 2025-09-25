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

	r.Post("/create", todoController.CreateTodoHandler)
	r.Put("/update", todoController.UpdateTodo)

	return r
}
