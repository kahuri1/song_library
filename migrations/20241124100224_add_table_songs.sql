-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS songs (
song_id SERIAL PRIMARY KEY,
title VARCHAR(255) NOT NULL,
group_id INT,
release_date DATE NOT NULL,
lyrics TEXT NOT NULL,
link VARCHAR(255) NOT NULL,
FOREIGN KEY (group_id) REFERENCES groups(group_id) ON DELETE CASCADE
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS "songs";
-- +goose StatementEnd
