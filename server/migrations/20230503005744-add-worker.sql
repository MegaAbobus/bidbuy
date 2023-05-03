
-- +migrate Up
CREATE TABLE "worker" (
  "worker_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "name" varchar,
  "pasport_series_number" varchar,
  "pasport_issued" varchar,
  "phone_number" varchar
);
-- +migrate Down
DROP TABLE "worker";