BEGIN;

INSERT INTO "user" VALUES ('00000000-0000-0000-0000-000000000000', 'ADMIN', 'admin@example.com', 'admin_google_id');
INSERT INTO "user" VALUES ('11111111-1111-1111-1111-111111111111', 'USER', 'user@example.com', 'user_google_id');

COMMIT;
