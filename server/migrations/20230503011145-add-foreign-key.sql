
-- +migrate Up
ALTER TABLE "bids" ADD CONSTRAINT "fk_bids_user" FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "order" ADD CONSTRAINT "fk_order_user"  FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "cart" ADD CONSTRAINT "fk_cart_user" FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "rewiews" ADD CONSTRAINT "fk_rewiews_user" FOREIGN KEY ("user_id") REFERENCES "user" ("user_id");

ALTER TABLE "appeal" ADD CONSTRAINT "fk_appeal_id" FOREIGN KEY ("appeal_id") REFERENCES "user" ("user_id");

ALTER TABLE "order_comp" ADD CONSTRAINT "fk_order_id" FOREIGN KEY ("order_id") REFERENCES "order" ("order_id");

ALTER TABLE "order_collector" ADD CONSTRAINT "fk_order_id" FOREIGN KEY ("order_id") REFERENCES "order" ("order_id");

ALTER TABLE "order_status_history" ADD CONSTRAINT "fk_order_id" FOREIGN KEY ("order_id") REFERENCES "order" ("order_id");

ALTER TABLE "cart_product" ADD CONSTRAINT "fk_cart_id" FOREIGN KEY ("cart_id") REFERENCES "cart" ("cart_id");

ALTER TABLE "auction" ADD CONSTRAINT "fk_auction_id" FOREIGN KEY ("auction_id") REFERENCES "product" ("product_id");

ALTER TABLE "rewiews" ADD CONSTRAINT "fk_rewiews_id" FOREIGN KEY ("product_id") REFERENCES "product" ("product_id");

ALTER TABLE "cart_product" ADD CONSTRAINT "fk_card_id" FOREIGN KEY ("product_id") REFERENCES "product" ("product_id");

ALTER TABLE "order_comp" ADD CONSTRAINT "fk_order_product" FOREIGN KEY ("product_id") REFERENCES "product" ("product_id");

ALTER TABLE "order_collect_history" ADD CONSTRAINT "fk_order_collector_id" FOREIGN KEY ("order_collector_id") REFERENCES "order_collector" ("order_collector_id");

ALTER TABLE "collector" ADD CONSTRAINT "fk_collector_id" FOREIGN KEY ("collector_id") REFERENCES "worker" ("worker_id");

ALTER TABLE "supporter" ADD CONSTRAINT "fk_supporter_id" FOREIGN KEY ("supporter_id") REFERENCES "worker" ("worker_id");

ALTER TABLE "order_collector" ADD CONSTRAINT "fk_order_collector_id" FOREIGN KEY ("collector_id") REFERENCES "collector" ("collector_id");

ALTER TABLE "appeal" ADD CONSTRAINT "fk_appeal_supporter_id" FOREIGN KEY ("appeal_id") REFERENCES "supporter" ("supporter_id");

ALTER TABLE "bids" ADD CONSTRAINT "fk_bids_auction" FOREIGN KEY ("auction_id") REFERENCES "auction" ("auction_id");
-- +migrate Down
ALTER TABLE "bids" DROP CONSTRAINT "fk_bids_user";

ALTER TABLE "order" DROP CONSTRAINT "fk_order_user";

ALTER TABLE "cart" DROP CONSTRAINT "fk_cart_user";

ALTER TABLE "rewiews" DROP CONSTRAINT "fk_rewiews_user";

ALTER TABLE "appeal" DROP CONSTRAINT "fk_appeal_id";

ALTER TABLE "order_comp" DROP CONSTRAINT "fk_order_id";

ALTER TABLE "order_collector" DROP CONSTRAINT "fk_order_id";

ALTER TABLE "order_status_history" DROP CONSTRAINT "fk_order_id";

ALTER TABLE "cart_product" DROP CONSTRAINT "fk_cart_id";

ALTER TABLE "auction" DROP CONSTRAINT "fk_auction_id";

ALTER TABLE "rewiews" DROP CONSTRAINT "fk_rewiews_id";

ALTER TABLE "cart_product" DROP CONSTRAINT "fk_card_id";

ALTER TABLE "order_comp" DROP CONSTRAINT "fk_order_product";

ALTER TABLE "order_collect_history" DROP CONSTRAINT "fk_order_collector_id";

ALTER TABLE "collector" DROP CONSTRAINT "fk_collector_id";

ALTER TABLE "supporter" DROP CONSTRAINT "fk_supporter_id";

ALTER TABLE "order_collector" DROP CONSTRAINT "fk_order_collector_id";

ALTER TABLE "appeal" DROP CONSTRAINT "fk_appeal_supporter_id";

ALTER TABLE "bids" DROP CONSTRAINT "fk_bids_auction";