BEGIN;

ALTER TABLE "shop" ADD COLUMN "owner_id" UUID;
ALTER TABLE "shop" ADD CONSTRAINT fk_shop_owner_id FOREIGN KEY (owner_id) REFERENCES "user" (id);

COMMIT;
