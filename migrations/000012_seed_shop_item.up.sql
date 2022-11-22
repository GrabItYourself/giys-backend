BEGIN;

INSERT INTO "shop_item" (shop_id, name, image, price) 
VALUES (1, 'ข้าวมันไก่ต้ม', NULL, 50) RETURNING "id";

INSERT INTO "shop_item" (shop_id, name, image, price)
VALUES (1, 'ข้าวมันไก่ทอด', NULL, 50) RETURNING "id";

INSERT INTO "shop_item" (shop_id, name, image, price)
VALUES (1, 'น้ำจิ้มรสเด็ด', NULL, 15) RETURNING "id";


INSERT INTO "shop_item" (shop_id, name, image, price)
VALUES (2, 'ก๋วยเตี๋ยวเรือ', NULL, 45) RETURNING "id";

INSERT INTO "shop_item" (shop_id, name, image, price)
VALUES (2, 'ก๋วยเตี๋ยวน้ำใส', NULL, 45) RETURNING "id";

INSERT INTO "shop_item" (shop_id, name, image, price)
VALUES (2, 'ก๋วยเตี๋ยวต้มยำ', NULL, 45) RETURNING "id";

INSERT INTO "shop_item" (shop_id, name, image, price)
VALUES (2, 'เกาเหลา', NULL, 30) RETURNING "id";


INSERT INTO "shop_item" (shop_id, name, image, price)
VALUES (3, 'หมี่หยกผัด', NULL, 45) RETURNING "id";

INSERT INTO "shop_item" (shop_id, name, image, price)
VALUES (3, 'ข้าวไข่เจียว', NULL, 40) RETURNING "id";


INSERT INTO "shop_item" (shop_id, name, image, price)
VALUES (4, 'ไอศกรีมรสโยเกิร์ต', NULL, 30) RETURNING "id";

COMMIT;
