
-- +migrate Up
CREATE TABLE "rewiews" (
  "rewiews" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "user_id" int,
  "product_id" int,
  "text" varchar,
  "rating" float
);
-- +migrate Down
DROP TABLE "rewiews";