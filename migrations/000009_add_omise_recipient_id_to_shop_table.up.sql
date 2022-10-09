BEGIN;

ALTER TABLE "shop" ADD COLUMN "omise_recipient_id" TEXT;

COMMIT;
