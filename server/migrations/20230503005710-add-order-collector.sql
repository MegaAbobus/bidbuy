
-- +migrate Up
CREATE TABLE "order_collector" (
  "order_collector_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "order_id" int,
  "collector_id" int,
  "status" varchar
);
-- +migrate Down
DROP TABLE "order_collector";