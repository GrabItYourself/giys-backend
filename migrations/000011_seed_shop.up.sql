BEGIN;

INSERT INTO "shop" (name, image, description, location, contact, omise_recipient_id) 
VALUES ('ข้าวมันไก่ป้าแดง', NULL, 'ข้าวมันไก่เจ้าดังที่ iCanteen', 'iCanteen ร้าน 12', '0812345678', 'recp_test_5tuutd6z5nsvxl729w2') RETURNING "id";

INSERT INTO "shop" (name, image, description, location, contact, omise_recipient_id) 
VALUES ('ก๋วยเตี๋ยวเรือ', NULL, NULL, 'โรงอาหารคณะอักษรฯ', '0987456123', 'recp_test_5tuuthgspiwqw7crjbb') RETURNING "id";

INSERT INTO "shop" (name, image, description, location, contact, omise_recipient_id) 
VALUES ('พี่ใหญ่ตามสั่ง', NULL, 'อยากได้ไรสั่งเลยจ้า พี่ทำให้ได้ทุกอย่าง ถ้าน้องกล้าสั่ง แต่ถ้าสั่งอะไรที่กินไม่ได้มาพี่ไม่รู้ด้วยนะ', 'iCanteen', NULL, 'recp_test_5tuutlu5u8sj6243nbc') RETURNING "id";

INSERT INTO "shop" (name, image, description, location, contact, omise_recipient_id) 
VALUES ('Parabola', NULL, 'ไอศกรีมในตำนาน', 'ลานเกียร์', NULL, 'recp_test_5tuutr6w2m7fkxhl7lw') RETURNING "id";

COMMIT;
