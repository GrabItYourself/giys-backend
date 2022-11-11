BEGIN;

INSERT INTO "shop_item" (shop_id, name, image, price) 
VALUES (1, 'ข้าวมันไก่ต้ม', NULL, 50) RETURNING "id";

INSERT INTO "shop_item" (shop_id, name, image, price)
VALUES (1, 'ข้าวมันไก่ทอด', NULL, 50) RETURNING "id";

INSERT INTO "shop_item" (shop_id, name, image, price)
VALUES (2, 'หมี่หยกผัด', NULL, 45) RETURNING "id";

INSERT INTO "shop_item" (shop_id, name, image, price)
VALUES (3, 'ข้าวไข่เจียว', NULL, 40) RETURNING "id";

COMMIT;
