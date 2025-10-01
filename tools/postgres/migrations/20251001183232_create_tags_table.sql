-- +goose Up
-- +goose StatementBegin
CREATE TABLE tags (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP tags;
-- +goose StatementEnd
