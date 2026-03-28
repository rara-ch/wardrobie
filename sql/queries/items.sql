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

-- name: ReadItems :many
SELECT * FROM items;

-- name: GetItemsByID :one
SELECT *
FROM items
WHERE id = $1;

-- name: DeleteItems :exec
DELETE FROM items;