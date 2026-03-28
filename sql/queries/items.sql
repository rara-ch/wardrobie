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

-- name: DeleteItems :exec
DELETE FROM items;