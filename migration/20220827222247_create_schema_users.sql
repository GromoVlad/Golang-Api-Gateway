-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA users;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA users;
-- +goose StatementEnd
