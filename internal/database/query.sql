-- name: CreateLevel :one
INSERT INTO levels (map) VALUES ($1)
RETURNING *;

-- name: GetLevel :one
SELECT * FROM levels
WHERE id = $1 LIMIT 1;

-- name: CreatePlayer :one
INSERT INTO players (
    level_id, start_x, start_y, curr_x, curr_y, hitpoints
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetPlayerByLevelIDAndLock :one
SELECT * FROM players WHERE level_id = $1 FOR UPDATE;

-- name: UpdatePlayer :execrows
UPDATE players
SET hitpoints = $2, curr_x = $3, curr_y = $4
WHERE level_id = $1;