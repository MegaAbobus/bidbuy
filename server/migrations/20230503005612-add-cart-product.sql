
-- +migrate Up
CREATE TABLE "cart_product" (
  "cart_product" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "cart_id" int,
  "product_id" int,
  "amount" int
);
-- +migrate Down
DROP TABLE "cart_product";