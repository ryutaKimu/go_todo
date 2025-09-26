package service

import (
	"context"
	"strconv"

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

func (s *TodoServiceImpl) FetchAllTodo(ctx context.Context) ([]*model.Todo, error) {
	return s.repo.FetchAllTodo(ctx)
}

func (s *TodoServiceImpl) FindTodoById(ctx context.Context, id string) (*model.Todo, error) {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	todo, err := s.repo.FindTodoById(ctx, userId)

	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TodoServiceImpl) CreateTodo(ctx context.Context, todo *model.Todo) error {
	return s.repo.CreateTodo(ctx, todo)
}

func (s *TodoServiceImpl) UpdateTodo(ctx context.Context, userId string, todo *model.Todo) error {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return err
	}
	return s.repo.UpdateTodo(ctx, id, todo)
}

func (s *TodoServiceImpl) DeleteTodo(ctx context.Context, userId string) error {
	id, err := strconv.Atoi(userId)

	if err != nil {
		return err
	}

	return s.repo.DeleteTodo(ctx, id)
}
