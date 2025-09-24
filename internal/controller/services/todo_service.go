package services

import (
	"context"

	"github.com/ryutaKimu/go_todo/internal/model"
)

type TodoService interface {
	CreateTodo(ctx context.Context, todo *model.Todo) error
}
