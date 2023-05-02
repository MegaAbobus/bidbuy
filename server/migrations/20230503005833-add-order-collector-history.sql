
-- +migrate Up
CREATE TABLE "order_collect_history" (
  "order_collect_history_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "order_collector_id" int,
  "status" varchar,
  "time" timestamp
);
-- +migrate Down
DROP TABLE "order_collect_history";