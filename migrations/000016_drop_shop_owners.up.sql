BEGIN;

ALTER TABLE "shop" DROP CONSTRAINT fk_shop_owner_id;
ALTER TABLE "shop" DROP COLUMN "owner_id";

COMMIT;
