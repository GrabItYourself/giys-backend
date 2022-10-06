BEGIN;

CREATE TABLE IF NOT EXISTS "payment_transaction" (
    id SERIAL PRIMARY KEY,
    shop_id INT NOT NULL,
    order_id INT NOT NULL,
    amount INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (shop_id, order_id) REFERENCES "order" (shop_id, id)
);

COMMIT;
