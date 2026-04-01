-- name: CreateItem :one
INSERT INTO items (id, created_at, updated_at, name, apparel, color, brand, material, category)
VALUES (
    gen_random_uuid(),
    now(),
    now(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetItems :many
SELECT * FROM items;

-- name: GetItemByName :one
SELECT *
FROM items
WHERE name = $1;

-- name: UpdateItem :one
UPDATE items
SET
    updated_at = now(),
    apparel = COALESCE($2, apparel),
    color = COALESCE($3, color),
    brand = COALESCE($4, brand),
    material = COALESCE($5, material),
    category = COALESCE($6, category)
WHERE name = $1
RETURNING *;

-- name: DeleteItems :exec
DELETE FROM items;

-- name: DeleteItemByName :one
DELETE FROM items
WHERE name = $1
RETURNING *;