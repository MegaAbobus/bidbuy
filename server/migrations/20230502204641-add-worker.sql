
-- +migrate Up
create table worker
(
    id                     serial
        constraint worker_pk
            primary key,
    name                   varchar(255),
    passport_series_number varchar(10) not null,
    passport_issued        varchar(255),
    phone_number           varchar(12)
);

-- +migrate Down
drop table worker;