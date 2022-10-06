BEGIN;

CREATE TYPE ORDER_STATUS AS ENUM ('IN_QUEUE', 'READY', 'COMPLETE', 'CANCELLED');
ALTER TABLE "order" ADD COLUMN "status" ORDER_STATUS;

COMMIT;
