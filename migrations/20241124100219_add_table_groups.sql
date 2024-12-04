-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS "groups" (
group_id SERIAL PRIMARY KEY,
name VARCHAR(255) NOT NULL
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS "groups"
-- +goose StatementEnd
