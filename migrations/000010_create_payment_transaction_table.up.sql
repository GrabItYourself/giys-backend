BEGIN;

CREATE TABLE IF NOT EXISTS "payment_transaction" (
    shop_id INT NOT NULL,
    order_id INT NOT NULL,
    amount INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (shop_id, order_id),
    FOREIGN KEY (shop_id, order_id) REFERENCES "order" (shop_id, id)
);

COMMIT;
