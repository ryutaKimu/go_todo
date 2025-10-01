-- +goose Up
-- +goose StatementBegin
CREATE TABLE todo_tags (
    todo_id BIGINT,
    tag_id BIGINT,
    PRIMARY KEY(todo_id, tag_id),
    FOREIGN KEY (todo_id) REFERENCES todos(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP todo_tags
-- +goose StatementEnd
