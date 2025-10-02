-- +goose Up
-- +goose StatementBegin
ALTER TABLE todo_tags
DROP CONSTRAINT todo_tags_todo_id_fkey,
ADD CONSTRAINT todo_tags_todo_id_fkey
FOREIGN KEY (todo_id) REFERENCES todos(id) ON DELETE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP todo_tags
-- +goose StatementEnd
