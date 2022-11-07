BEGIN;

ALTER TABLE "payment_method" ADD COLUMN "last_four_digits" TEXT;

COMMIT;
