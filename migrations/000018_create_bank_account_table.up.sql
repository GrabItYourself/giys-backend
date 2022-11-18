BEGIN;

CREATE TABLE IF NOT EXISTS "bank_account" (
    shop_id INT PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    brand TEXT NOT NULL,
    number TEXT NOT NULL,
    FOREIGN KEY (shop_id) REFERENCES "shop" (id) ON DELETE CASCADE
);

COMMIT;
