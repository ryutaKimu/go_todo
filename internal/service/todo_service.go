package service

import (
	"context"

	"github.com/ryutaKimu/go_todo/internal/controller/services"
	"github.com/ryutaKimu/go_todo/internal/model"
	"github.com/ryutaKimu/go_todo/internal/repository"
)

type TodoServiceImpl struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) services.TodoService {
	return &TodoServiceImpl{repo: repo}
}

func (s *TodoServiceImpl) CreateTodo(ctx context.Context, todo *model.Todo) error {
	return s.repo.CreateTodo(ctx, todo)
}
