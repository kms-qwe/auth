-- +goose Up
create table role (
    id integer primary key, 
    role_name text not null
);

INSERT INTO role (id, role_name)
VALUES
    (0, 'user'),
    (1, 'admin');
-- +goose Down
drop table role;
