package todo

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/ryutaKimu/go_todo/internal/model"
	"github.com/ryutaKimu/go_todo/internal/repository"
)

type TodoRepositoryImpl struct {
	db   *sql.DB
	goqu goqu.DialectWrapper
}

func NewTodoRepository(db *sql.DB) repository.TodoRepository {
	return &TodoRepositoryImpl{
		db:   db,
		goqu: goqu.Dialect("postgres"),
	}
}

func (r *TodoRepositoryImpl) CreateTodo(ctx context.Context, todo *model.Todo) error {
	record := goqu.Record{
		"title":        todo.Title,
		"is_completed": todo.IsCompleted,
	}

	query, args, err := r.goqu.Insert("todos").Rows(record).ToSQL()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)

	if err != nil {
		return err
	}

	return nil
}
