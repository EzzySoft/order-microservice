-- 0001_create_orders_table.up.sql

CREATE TABLE orders (
    uid VARCHAR PRIMARY KEY,
    track_number VARCHAR,
    entry VARCHAR,
    internal_signature VARCHAR,
    customer_id VARCHAR,
    delivery_service VARCHAR,
    shardkey VARCHAR,
    sm_id INTEGER,
    date_created TIMESTAMP,
    oof_shard VARCHAR
);

CREATE TABLE deliveries (
    order_uid VARCHAR PRIMARY KEY REFERENCES orders(uid),
    name VARCHAR,
    phone VARCHAR,
    zip VARCHAR,
    city VARCHAR,
    address VARCHAR,
    region VARCHAR,
    email VARCHAR
);

CREATE TABLE payments (
    order_uid VARCHAR PRIMARY KEY REFERENCES orders(uid),
    transaction VARCHAR,
    request_id VARCHAR,
    currency VARCHAR,
    provider VARCHAR,
    amount INTEGER,
    payment_dt BIGINT,
    bank VARCHAR,
    delivery_cost INTEGER,
    goods_total INTEGER,
    custom_fee INTEGER
);

CREATE TABLE items (
    order_uid VARCHAR REFERENCES orders(uid),
    chrt_id INTEGER,
    track_number VARCHAR,
    price INTEGER,
    rid VARCHAR,
    name VARCHAR,
    sale INTEGER,
    size VARCHAR,
    total_price INTEGER,
    nm_id INTEGER,
    brand VARCHAR,
    status INTEGER,
    PRIMARY KEY (order_uid, chrt_id)
);
