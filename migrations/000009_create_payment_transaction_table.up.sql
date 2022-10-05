BEGIN;

CREATE TABLE IF NOT EXISTS "payment_method" (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    amount INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (order_id) REFERENCES "order" (id)
);

COMMIT;
