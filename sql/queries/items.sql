-- name: CreateItem :one
INSERT INTO items (id, created_at, updated_at, type, color, brand, material, category)
VALUES (
    gen_random_uuid(),
    now(),
    now(),
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetItems :many
SELECT * FROM items;

-- name: GetItemByID :one
SELECT *
FROM items
WHERE id = $1;

-- name: UpdateItem :one
UPDATE items
SET
    updated_at = now(),
    type = COALESCE($2, type),
    color = COALESCE($3, color),
    brand = COALESCE($4, brand),
    material = COALESCE($5, material),
    category = COALESCE($6, category)
WHERE id = $1
RETURNING *;

-- name: DeleteItems :exec
DELETE FROM items;