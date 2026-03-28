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

-- name: DeleteItems :exec
DELETE FROM items;