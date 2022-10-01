BEGIN;

CREATE TYPE role AS ENUM ('admin', 'user');
CREATE TABLE IF NOT EXISTS "users" (
    id TEXT PRIMARY KEY,
    role role NOT NULL,
    email TEXT NOT NULL
);

COMMIT;
