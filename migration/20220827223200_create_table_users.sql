-- +goose Up
-- +goose StatementBegin
CREATE TABLE users.users (
    user_id serial primary key,
    name varchar(255) not null,
    role_id smallint not null,
    phone varchar(255) null,
    password varchar(512) null,
    email varchar(255) null,
    horeca_id integer null,
    password_recovery_url varchar(1024) null,
    messenger varchar(255) null,
    created_at timestamp null,
    updated_at timestamp null,
    constraint users_email_unique unique (email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users.users
-- +goose StatementEnd
