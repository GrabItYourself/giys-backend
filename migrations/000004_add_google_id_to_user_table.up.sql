BEGIN;

ALTER TABLE "user" ADD COLUMN "google_id" TEXT;

COMMIT;
