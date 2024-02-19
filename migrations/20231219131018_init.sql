-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS stocks (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name varchar NOT NULL,
    is_available boolean NOT NULL DEFAULT TRUE,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS items (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name varchar NOT NULL,
    size varchar NOT NULL,
    quantity bigint NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS reserved_items (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    stock_id uuid NOT NULL,
    item_id uuid NOT NULL,
    quantity bigint NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    CONSTRAINT items_stocks_stocks_fkey
        FOREIGN KEY (stock_id) REFERENCES stocks (id),
    CONSTRAINT reserved_items_items_fkey
        FOREIGN KEY (item_id) REFERENCES items (id),
    UNIQUE(stock_id, item_id),
    PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS items_stocks (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    stock_id uuid NOT NULL,
    item_id uuid NOT NULL,
    quantity bigint NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    CONSTRAINT items_stocks_stocks_fkey
        FOREIGN KEY (stock_id) REFERENCES stocks (id),
    CONSTRAINT items_stocks_items_fkey
        FOREIGN KEY (item_id) REFERENCES items (id),
    UNIQUE(stock_id, item_id),
    PRIMARY KEY (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reserved_items;
DROP TABLE IF EXISTS items_stocks;
DROP TABLE IF EXISTS stocks;
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
