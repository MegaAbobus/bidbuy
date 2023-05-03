
-- +migrate Up
CREATE TABLE "cart" (
  "cart_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "user_id" int
);
-- +migrate Down
DROP TABLE "cart";