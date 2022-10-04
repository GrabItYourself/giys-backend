BEGIN;

CREATE TABLE IF NOT EXISTS "payment_method" (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    omise_card_id TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user" (id)
);

COMMIT;
