BEGIN;

ALTER TABLE "user" 
ADD COLUMN "default_payment_method_id" INT;

ALTER TABLE "user" 
ADD CONSTRAINT fk_default_payment_method
FOREIGN KEY (default_payment_method_id) 
REFERENCES "payment_method" (id)
ON DELETE CASCADE;

COMMIT;
