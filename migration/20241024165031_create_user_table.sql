-- +goose Up
create table userV1.user (
    id integer generated always as identity primary key,
    name text not null,
    email text not null,
    password text not null,
    role integer not null default 0,
    created_at  timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
DROP table userV1.user;