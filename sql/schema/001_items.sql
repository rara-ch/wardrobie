-- +goose Up
CREATE TABLE items (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    color VARCHAR NOT NULL,
    type VARCHAR NOT NULL
);

-- +goose Down
DROP TABLE items;