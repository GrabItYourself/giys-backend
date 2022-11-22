BEGIN;

INSERT INTO "user" VALUES ('00000000-0000-0000-0000-000000000000', 'ADMIN', 'admin@example.com', 'admin_google_id');
INSERT INTO "user" VALUES ('11111111-1111-1111-1111-111111111111', 'USER', 'alice@example.com', 'alice_google_id');
INSERT INTO "user" VALUES ('22222222-2222-2222-2222-222222222222', 'USER', 'bob@example.com', 'bob_google_id');
INSERT INTO "user" VALUES ('33333333-3333-3333-3333-333333333333', 'USER', 'clara@example.com', 'clara_google_id');

COMMIT;
