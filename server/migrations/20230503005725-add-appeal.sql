
-- +migrate Up
CREATE TABLE "appeal" (
  "appeal_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "user_id" int,
  "supporter" int,
  "message" varchar
);
-- +migrate Down
DROP TABLE "appeal";