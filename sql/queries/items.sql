-- name: CreateItem :one
INSERT INTO items (id, created_at, updated_at, color, type)
VALUES (
    gen_random_uuid(),
    now(),
    now(),
    $1,
    $2
)
RETURNING *;

-- name: DeleteItems :exec
DELETE FROM items;