BEGIN;

DELETE FROM "bank_account" WHERE shop_id=1;
DELETE FROM "bank_account" WHERE shop_id=2;
DELETE FROM "bank_account" WHERE shop_id=3;
DELETE FROM "bank_account" WHERE shop_id=4;

COMMIT;
