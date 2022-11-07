BEGIN;

ALTER TABLE "payment_method" DROP COLUMN "last_four_digits";

COMMIT;
