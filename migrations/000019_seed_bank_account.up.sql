BEGIN;

INSERT INTO "bank_account" (shop_id, name, type, brand, number) 
VALUES (1, 'Alice1', 'individual', 'scb', '123456789');

INSERT INTO "bank_account" (shop_id, name, type, brand, number) 
VALUES (2, 'Alice2', 'individual', 'kbank', '12345678');

INSERT INTO "bank_account" (shop_id, name, type, brand, number) 
VALUES (3, 'Bob1', 'individual', 'kbank', '12345678');

INSERT INTO "bank_account" (shop_id, name, type, brand, number) 
VALUES (4, 'Bob2', 'individual', 'scb', '123456789');

COMMIT;