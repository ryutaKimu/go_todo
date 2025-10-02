package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"

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
	todos, err := s.repo.FetchAllTodo(ctx)

	if len(todos) == 0 {
		return nil, fmt.Errorf("todoが存在しません: %w", sql.ErrNoRows)
	}

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *TodoServiceImpl) FindTodoById(ctx context.Context, id string) (*model.Todo, error) {
	todoId, err := strconv.Atoi(id)

	if err != nil {
		return nil, fmt.Errorf("不正なパラムが検知されました。%w", err)
	}

	todo, err := s.repo.FindTodoById(ctx, todoId)

	if errors.Is(err, sql.ErrNoRows) {
		err := fmt.Errorf("Todoは存在しません:%w", sql.ErrNoRows)
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TodoServiceImpl) CreateTodo(ctx context.Context, todo *model.Todo) error {
	const MAX_TITLE_LENGTH = 20

	if utf8.RuneCountInString(todo.Title) > MAX_TITLE_LENGTH {
		return fmt.Errorf("タイトルは%d文字で入力してください", MAX_TITLE_LENGTH)
	}
	return s.repo.CreateTodo(ctx, todo)
}

func (s *TodoServiceImpl) UpdateTodo(ctx context.Context, id string, todo *model.Todo) error {
	todoId, err := strconv.Atoi(id)

	if err != nil {
		return fmt.Errorf("不正なパラムが検知されました。%w", err)
	}

	err = s.repo.UpdateTodo(ctx, todoId, todo)
	if errors.Is(err, sql.ErrNoRows) {
		err := fmt.Errorf("Todoは存在しません:%w", sql.ErrNoRows)
		return err
	}

	err = s.repo.UpdateTodoTags(ctx, todoId, todo)

	if err != nil {
		return err
	}

	return nil
}

func (s *TodoServiceImpl) DeleteTodo(ctx context.Context, todoId string) error {
	id, err := strconv.Atoi(todoId)
	if err != nil {
		return fmt.Errorf("不正なパラムが検知されました。%w", err)
	}

	err = s.repo.DeleteTodo(ctx, id)

	if errors.Is(err, sql.ErrNoRows) {
		err := fmt.Errorf("Todoは存在しません:%w", sql.ErrNoRows)
		return err
	}

	return nil
}
