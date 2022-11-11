BEGIN;

INSERT INTO "shop" (name, image, description, location, contact, omise_recipient_id) 
VALUES ('Nana', NULL, 'It is a long established fact that a reader will be distracted', NULL, NULL, NULL) RETURNING "id";

INSERT INTO "shop" (name, image, description, location, contact, omise_recipient_id) 
VALUES ('Nima', NULL, 'It is a long established fact that a reader will be distracted', NULL, NULL, NULL) RETURNING "id";

INSERT INTO "shop" (name, image, description, location, contact, omise_recipient_id) 
VALUES ('Namu', NULL, 'It is a long established fact that a reader will be distracted', NULL, NULL, NULL) RETURNING "id";

INSERT INTO "shop" (name, image, description, location, contact, omise_recipient_id) 
VALUES ('Ninni', NULL, 'It is a long established fact that a reader will be distracted', NULL, NULL, NULL) RETURNING "id";

COMMIT;
