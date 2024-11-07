-- +goose Up
create table users (
    id integer generated always as identity primary key,
    name text not null,
    email text not null,
    password text not null,
    role integer not null default 0,
    created_at  timestamp not null default now(),
    updated_at timestamp,

    constraint fk_role foreign key (role) references role(id) 
);

-- +goose Down
DROP table users;