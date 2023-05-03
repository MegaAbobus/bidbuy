
-- +migrate Up
CREATE TABLE "order_status_history" (
  "order_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "status" varchar,
  "time" timestamp
);
-- +migrate Down
DROP TABLE "order_status_history";