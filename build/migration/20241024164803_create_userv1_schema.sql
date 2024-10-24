-- +goose Up
create schema userV1;

-- +goose Down
drop schema userV1 restrict;