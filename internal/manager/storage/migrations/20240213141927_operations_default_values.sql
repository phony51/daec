-- +goose Up
-- +goose StatementBegin
INSERT INTO operations (name, latency) VALUES ('+', 200);
INSERT INTO operations (name, latency) VALUES ('-', 200);
INSERT INTO operations (name, latency) VALUES ('*', 200);
INSERT INTO operations (name, latency) VALUES ('/', 200);
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM operations WHERE name = '+';
DELETE FROM operations WHERE name = '-';
DELETE FROM operations WHERE name = '*';
DELETE FROM operations WHERE name = '/';
SELECT 'down SQL query';
-- +goose StatementEnd
