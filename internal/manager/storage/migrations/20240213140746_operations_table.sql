-- +goose Up
-- +goose StatementBegin
create table operations
(
    name    TEXT,
    latency INTEGER
);

SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table operations;
SELECT 'down SQL query';
-- +goose StatementEnd
