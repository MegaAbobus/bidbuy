
-- +migrate Up
CREATE TABLE "supporter" (
  "supporter_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "worker_id" int
);
-- +migrate Down
DROP TABLE "supporter";