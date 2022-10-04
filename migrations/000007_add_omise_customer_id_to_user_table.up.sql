BEGIN;

ALTER TABLE "user" ADD COLUMN "omise_customer_id" TEXT;

COMMIT;