BEGIN;

DROP TRIGGER "trig_order_item_id" ON "order_item";
DROP TRIGGER "trig_order_id" ON "order";
DROP TABLE IF EXISTS "order_item";
DROP TABLE IF EXISTS "order";
DROP FUNCTION "fn_order_item_id"();
DROP FUNCTION "fn_order_id"();

COMMIT;
