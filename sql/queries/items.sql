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
    name = COALESCE($2, name),
    apparel = COALESCE($3, apparel),
    color = COALESCE($4, color),
    brand = COALESCE($5, brand),
    material = COALESCE($6, material),
    category = COALESCE($7, category)
WHERE name = $1
RETURNING *;

-- name: DeleteItems :exec
DELETE FROM items;

-- name: DeleteItemByName :one
DELETE FROM items
WHERE name = $1
RETURNING *;