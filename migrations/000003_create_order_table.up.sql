BEGIN;

-- Create order table
CREATE TABLE IF NOT EXISTS "orders" (
    id INT NOT NULL,
    shop_id INT NOT NULL,
    user_id TEXT NOT NULL,
    PRIMARY KEY (shop_id, id),
    FOREIGN KEY (shop_id) REFERENCES shop(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- The function that can automatically generate the next id for the order table
CREATE OR REPLACE FUNCTION "fn_order_id"()
RETURNS "pg_catalog"."trigger" AS $BODY$ 
	BEGIN
		NEW.id = (SELECT COUNT(*)+1 FROM "orders" WHERE shop_id=NEW.shop_id);
		return NEW;
	END;
$BODY$
LANGUAGE plpgsql VOLATILE;

-- Trigger the function when a new row is inserted into the "order" table
CREATE TRIGGER "trig_order_pk" BEFORE INSERT 
  ON "orders"
  FOR EACH ROW
  EXECUTE PROCEDURE fn_order_id();

-- Create order_item table
CREATE TABLE IF NOT EXISTS "order_item" (
    id INT NOT NULL,
    order_id INT NOT NULL,
    shop_id INT NOT NULL,
    PRIMARY KEY (shop_id, order_id, id),
    FOREIGN KEY (shop_id, order_id) REFERENCES orders(shop_id, id)
);

-- The function that can automatically generate the next id for the order_item table
CREATE OR REPLACE FUNCTION "fn_order_item_id"()
RETURNS "pg_catalog"."trigger" AS $BODY$ 
	BEGIN
		NEW.id = (SELECT COUNT(*)+1 FROM "order_item" WHERE shop_id=NEW.shop_id AND order_id=NEW.order_id);
		return NEW;
	END;
$BODY$
LANGUAGE plpgsql VOLATILE;

-- Trigger the function when a new row is inserted into the "order_item" table
CREATE TRIGGER "trig_order_item_pk" BEFORE INSERT 
  ON "order_item"
  FOR EACH ROW
  EXECUTE PROCEDURE fn_order_item_id();

COMMIT;
