BEGIN;

CREATE TABLE IF NOT EXISTS "shop_owner" (
    shop_id INT NOT NULL,
    user_id UUID NOT NULL,

    PRIMARY KEY (shop_id, user_id),
    FOREIGN KEY (shop_id) REFERENCES "shop" (id),
    FOREIGN KEY (user_id) REFERENCES "user" (id)
);

COMMIT;
