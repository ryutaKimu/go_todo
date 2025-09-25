package repository

import (
	"context"

	"github.com/ryutaKimu/go_todo/internal/model"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *model.Todo) error
	UpdateTodo(ctx context.Context, todo *model.Todo) error
}
