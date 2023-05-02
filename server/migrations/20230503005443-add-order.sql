
-- +migrate Up
CREATE TABLE "order" (
  "order_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "user_id" int,
  "address" varchar,
  "price" float,
  "status" varchar
);
-- +migrate Down
DROP TABLE "order";