package repository

import (
	"context"

	"github.com/ryutaKimu/go_todo/internal/model"
)

type TodoRepository interface {
	FetchAllTodo(ctx context.Context) ([]*model.Todo, error)
	CreateTodo(ctx context.Context, todo *model.Todo) error
	UpdateTodo(ctx context.Context, userId int, todo *model.Todo) error
	DeleteTodo(ctx context.Context, id int) error
}
