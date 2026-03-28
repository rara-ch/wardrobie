-- +goose Up
AlTER TABLE items 
ADD COLUMN brand VARCHAR(100),
ADD COLUMN material VARCHAR(100),
ADD COLUMN category VARCHAR(100);

-- +goose Down
ALTER TABLE items
DROP COLUMN brand,
DROP COLUMN material,
DROP COLUMN category;