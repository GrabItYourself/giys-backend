BEGIN;

DELETE FROM "shop_owner" WHERE shop_id=1;
DELETE FROM "shop_owner" WHERE shop_id=2;
DELETE FROM "shop_owner" WHERE shop_id=3;
DELETE FROM "shop_owner" WHERE shop_id=4;

COMMIT;