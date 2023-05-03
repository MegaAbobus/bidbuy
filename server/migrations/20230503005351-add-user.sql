
-- +migrate Up
CREATE TABLE "user" (
  "user_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "phone_number" varchar,
  "amount_of_ransom" float,
  "email" varchar
);

-- +migrate Down
DROP TABLE "user";
