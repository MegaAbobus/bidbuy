
-- +migrate Up
CREATE TABLE "order_comp" (
  "order_comp" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "order_id" int,
  "product_id" int,
  "amount" int
);
-- +migrate Down
DROP TABLE "order_comp";