package services

import (
	"context"

	"github.com/ryutaKimu/go_todo/internal/model"
)

type TodoService interface {
	FetchAllTodo(ctx context.Context) ([]*model.Todo, error)
	FindTodoById(ctx context.Context, userId string) (*model.Todo, error)
	CreateTodo(ctx context.Context, todo *model.Todo) error
	UpdateTodo(ctx context.Context, userId string, todo *model.Todo) error
	DeleteTodo(ctx context.Context, userId string) error
}
