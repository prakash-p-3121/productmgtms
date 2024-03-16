CREATE TABLE products (
    id VARBINARY(2000) NOT NULL PRIMARY KEY,
    id_bit_count BIGINT NOT NULL,
    name VARCHAR(10000) NOT NULL,
    description TEXT NOT NULL,
    currency VARCHAR(10) NOT NULL,
    cost_price DOUBLE NOT NULL,
    category_id VARBINARY(2000) NOT NULL,
    media_type INT UNSIGNED NOT NULL,
    media_path TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX products_cost_price (cost_price),
    INDEX products_category_id (category_id)
);