BEGIN;

DELETE FROM "user" WHERE id = '00000000-0000-0000-0000-000000000000';
DELETE FROM "user" WHERE id = '11111111-1111-1111-1111-111111111111';

COMMIT;
