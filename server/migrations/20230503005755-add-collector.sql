
-- +migrate Up
CREATE TABLE "collector" (
  "collector_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "worker_id" int
);
-- +migrate Down
DROP TABLE "collector";