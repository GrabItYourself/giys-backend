BEGIN;

DELETE FROM "shop" WHERE id=1;
DELETE FROM "shop" WHERE id=2;
DELETE FROM "shop" WHERE id=3;
DELETE FROM "shop" WHERE id=4;

COMMIT;