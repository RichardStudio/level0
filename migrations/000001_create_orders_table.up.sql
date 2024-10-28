CREATE TABLE orders (
    id VARCHAR(255) PRIMARY KEY,
    order_uid VARCHAR(255),
    track_number VARCHAR(255),
    entry VARCHAR(255),
    delivery JSONB,
    payment JSONB,
    items JSONB,
    locale VARCHAR(255),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(255),
    delivery_service VARCHAR(255),
    shardkey VARCHAR(255),
    sm_id INTEGER,
    date_created VARCHAR(255),
    oof_shard VARCHAR(255)
);
