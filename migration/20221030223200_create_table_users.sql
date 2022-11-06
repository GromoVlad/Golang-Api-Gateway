-- +goose Up
-- +goose StatementBegin
CREATE TABLE users.users (
    user_id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    role_id SMALLINT NOT NULL,
    password VARCHAR(512) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NULL,
    venue_id INTEGER NULL,
    password_recovery_url VARCHAR(1024) NULL,
    messenger VARCHAR(255) NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    CONSTRAINT users_email_unique UNIQUE (email),
    FOREIGN KEY (role_id) REFERENCES auth.roles (role_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users.users
-- +goose StatementEnd
