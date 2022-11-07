BEGIN;

ALTER TABLE "user" DROP CONSTRAINT fk_default_payment_method;
ALTER TABLE "user" DROP COLUMN "default_payment_method_id";

COMMIT;
