-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA auth;

CREATE TABLE auth.roles (
    role_id serial PRIMARY KEY,
    code VARCHAR NOT NULL,
    description text DEFAULT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);

comment ON TABLE auth.roles IS 'Роли пользователей приложения';

INSERT INTO auth.roles (role_id, code, description) VALUES
    (1, 'MANAGER', 'Менеджер заведения'),
    (2, 'WAITER', 'Официант'),
    (3, 'SUPPORT', 'Техническая поддержка')
ON conflict do nothing;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA auth CASCADE;
-- +goose StatementEnd
