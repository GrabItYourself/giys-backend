BEGIN;

CREATE TYPE ORDER_STATUS AS ENUM ('IN_QUEUE', 'READY', 'COMPLETED', 'CANCELED');
ALTER TABLE "order" ADD COLUMN "status" ORDER_STATUS;

COMMIT;
