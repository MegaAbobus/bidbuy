
-- +migrate Up
create table product
(
    id           serial
        constraint product_pk
            primary key,
    description  varchar(500),
    body         varchar(1000),
    is_available bit default 0::bit not null,
    price        real,
    rating       real
);

-- +migrate Down
drop table product;