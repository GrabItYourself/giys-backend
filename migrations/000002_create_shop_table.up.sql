BEGIN;

CREATE TABLE IF NOT EXISTS "shop" (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    image TEXT,
    description TEXT,
    location TEXT,
    contact TEXT 
);

CREATE TABLE IF NOT EXISTS "shop_item" (
    id SERIAL,
    shop_id INT NOT NULL,
    name TEXT NOT NULL,
    image TEXT,
    price INT NOT NULL,

    FOREIGN KEY (shop_id) REFERENCES "shop" (id) ON DELETE CASCADE,
    PRIMARY KEY (shop_id, id)
);

COMMIT;
