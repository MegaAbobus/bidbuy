
-- +migrate Up
CREATE TABLE "bids" (
  "bids" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "user_id" int,
  "auction_id" int,
  "bid" float,
  "time" timestamp
);
-- +migrate Down
DROP TABLE "bids";