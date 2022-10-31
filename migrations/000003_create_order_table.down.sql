BEGIN;

DROP TRIGGER "trig_order_id" ON "order";
DROP TABLE IF EXISTS "order_item";
DROP TABLE IF EXISTS "order";
DROP FUNCTION "fn_order_id"();

COMMIT;
