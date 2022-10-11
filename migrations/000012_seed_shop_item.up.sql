BEGIN;

INSERT INTO "shop_item" (id, shop_id, name, image, price) 
VALUES (1, 1, 'ข้าวมันไก่ต้ม', NULL, 50);

INSERT INTO "shop_item" (id, shop_id, name, image, price)
VALUES (2, 1, 'ข้าวมันไก่ทอด', NULL, 50);

INSERT INTO "shop_item" (id, shop_id, name, image, price)
VALUES (3, 2, 'หมี่หยกผัด', NULL, 45);

INSERT INTO "shop_item" (id, shop_id, name, image, price)
VALUES (4, 3, 'ข้าวไข่เจียว', NULL, 40);

COMMIT;
