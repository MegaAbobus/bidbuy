
-- +migrate Up
CREATE TABLE "auction" (
  "auction_id" int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  "product_id" int,
  "start_time" date,
  "end_time" date,
  "start_bid" float
);
-- +migrate Down
DROP TABLE "auction";