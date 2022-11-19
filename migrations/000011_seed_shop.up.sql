BEGIN;

INSERT INTO "shop" (name, image, description, location, contact, omise_recipient_id) 
VALUES ('Nana', NULL, 'It is a long established fact that a reader will be distracted', NULL, NULL, 'recp_test_5tuutd6z5nsvxl729w2') RETURNING "id";

INSERT INTO "shop" (name, image, description, location, contact, omise_recipient_id) 
VALUES ('Nima', NULL, 'It is a long established fact that a reader will be distracted', NULL, NULL, 'recp_test_5tuuthgspiwqw7crjbb') RETURNING "id";

INSERT INTO "shop" (name, image, description, location, contact, omise_recipient_id) 
VALUES ('Namu', NULL, 'It is a long established fact that a reader will be distracted', NULL, NULL, 'recp_test_5tuutlu5u8sj6243nbc') RETURNING "id";

INSERT INTO "shop" (name, image, description, location, contact, omise_recipient_id) 
VALUES ('Ninni', NULL, 'It is a long established fact that a reader will be distracted', NULL, NULL, 'recp_test_5tuutr6w2m7fkxhl7lw') RETURNING "id";

COMMIT;
