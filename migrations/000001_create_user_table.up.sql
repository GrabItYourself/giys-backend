BEGIN;

CREATE TYPE role AS ENUM ('admin', 'user');
CREATE TABLE IF NOT EXISTS "user" (
    id TEXT PRIMARY KEY,
    role role NOT NULL DEFAULT 'user',
    email TEXT NOT NULL
);

COMMIT;
