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

func (r *TodoRepositoryImpl) FetchAllTodo(ctx context.Context) ([]*model.Todo, error) {
	query, args, err := r.goqu.
		From("todos").
		Select("id", "title", "is_completed").
		ToSQL()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, err
	}

	defer rows.Close()

	var todos []*model.Todo

	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.Id, &t.Title, &t.IsCompleted); err != nil {
			return nil, err
		}
		todos = append(todos, &t)
	}

	return todos, nil
}

func (r *TodoRepositoryImpl) FindTodoById(ctx context.Context, userId int) (*model.Todo, error) {
	query, args, err := r.goqu.
		From("todos").
		Select("id", "title", "is_completed").
		Where(goqu.Ex{
			"id": userId,
		}).
		ToSQL()

	if err != nil {
		return nil, err
	}

	todo := &model.Todo{}
	err = r.db.QueryRowContext(ctx, query, args...).
		Scan(&todo.Id, &todo.Title, &todo.IsCompleted)

	if err == sql.ErrNoRows {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *TodoRepositoryImpl) CreateTodo(ctx context.Context, todo *model.Todo) error {
	record := goqu.Record{
		"title":        todo.Title,
		"is_completed": todo.IsCompleted,
	}

	query, args, err := r.goqu.Insert("todos").Rows(record).Returning("id").ToSQL()

	if err != nil {
		return err
	}

	var todoID int
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&todoID)

	if err != nil {
		return err
	}

	for _, tagId := range todo.TagIds {
		record := goqu.Record{
			"todo_id": todoID,
			"tag_id":  tagId,
		}
		query, args, err := r.goqu.Insert("todo_tags").Rows(record).ToSQL()

		if err != nil {
			return err
		}
		_, err = r.db.Exec(query, args...)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *TodoRepositoryImpl) UpdateTodo(ctx context.Context, userId int, todo *model.Todo) error {
	query, args, err := r.goqu.Update("todos").
		Set(goqu.Record{
			"title":        todo.Title,
			"is_completed": todo.IsCompleted,
			"updated_at":   goqu.L("NOW()"),
		}).
		Where(goqu.Ex{
			"id": userId,
		}).
		ToSQL()

	if err != nil {
		return err
	}

	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil

}

func (r *TodoRepositoryImpl) DeleteTodo(ctx context.Context, userId int) error {
	query, _, err := r.goqu.Delete("todos").
		Where(goqu.Ex{"id": userId}).
		ToSQL()

	if err != nil {
		return err
	}

	res, err := r.db.ExecContext(ctx, query)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
