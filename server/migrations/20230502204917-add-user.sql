
-- +migrate Up
create table "user"
(
    id              integer generated always as identity
        constraint user_pk
            primary key,
    phone_number     varchar(12),
    amount_of_ransom real,
    email            varchar(255)
);
-- +migrate Down
drop table "user";