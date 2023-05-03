
-- +migrate Up
CREATE TABLE "product" (
  "product_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "description" varchar,
  "body" varchar,
  "is_alailable" bit,
  "price" float,
  "rating" float
);

-- +migrate Down
DROP TABLE "product";