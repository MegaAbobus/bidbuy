-- +migrate Up
create table "order"
(
    id      serial
        constraint order_pk
            primary key,
    user_id serial,
    address varchar(255),
    price   real not null,
    status  varchar(255)
);

-- +migrate Down
drop table "order";