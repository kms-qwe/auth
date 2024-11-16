-- +goose Up
create table role (
    id integer primary key, 
    role_name text not null
);

INSERT INTO role (id, role_name)
VALUES
    (0, 'unknown'),
    (1, 'user'),
    (2, 'admin');
-- +goose Down
drop table role;
