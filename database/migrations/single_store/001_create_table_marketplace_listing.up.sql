CREATE ROWSTORE TABLE marketplace_listings (
    id VARBINARY(2000) NOT NULL,
    id_bit_count BIGINT UNSIGNED NOT NULL,
    product_id VARBINARY(2000) NOT NULL,
    seller_id VARBINARY(2000) NOT NULL,
    currency VARCHAR(10) NOT NULL,
    selling_price DOUBLE NOT NULL,
    stock_count BIGINT UNSIGNED DEFAULT NULL,
    return_policy INT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    SHARD KEY (product_id, seller_id),
    PRIMARY KEY (id, product_id, seller_id),
    UNIQUE(product_id, seller_id)
);
