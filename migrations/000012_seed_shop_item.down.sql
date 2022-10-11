BEGIN;

DELETE FROM "shop_item" WHERE id=1;
DELETE FROM "shop_item" WHERE id=2;
DELETE FROM "shop_item" WHERE id=3;
DELETE FROM "shop_item" WHERE id=4;

COMMIT;
