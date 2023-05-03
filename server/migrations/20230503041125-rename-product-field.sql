
-- +migrate Up
alter table "product"
    rename column is_alailable to is_available;
-- +migrate Down
alter table "product"
    rename column is_available to is_alilable;