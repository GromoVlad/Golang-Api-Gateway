-- +goose Up
-- +goose StatementBegin
CREATE TABLE auth.jwt_tokens
(
    jwt_token_id  serial PRIMARY KEY,
    refresh_token VARCHAR                 NOT NULL,
    user_id       INTEGER UNIQUE          NOT NULL,
    created_at    TIMESTAMP DEFAULT now() NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users.users (user_id)
);

comment ON TABLE auth.jwt_tokens IS 'Refresh токены пользователей приложения';

CREATE UNIQUE INDEX refresh_token_index ON auth.jwt_tokens (refresh_token);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE auth.jwt_tokens;
-- +goose StatementEnd
