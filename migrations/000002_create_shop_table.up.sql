BEGIN;

CREATE TABLE IF NOT EXISTS "shop" (
    id SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS "shop_item" (
    id INT NOT NULL,
    shop_id INT NOT NULL,

    FOREIGN KEY (shop_id) REFERENCES "shop" (id),
    PRIMARY KEY (shop_id, id)
);

COMMIT;
