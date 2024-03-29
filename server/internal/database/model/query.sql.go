// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package model

import (
	"context"
	"encoding/json"
)

const createLevel = `-- name: CreateLevel :one
INSERT INTO levels (map) VALUES ($1)
RETURNING id, map
`

func (q *Queries) CreateLevel(ctx context.Context, map_ json.RawMessage) (Level, error) {
	row := q.db.QueryRowContext(ctx, createLevel, map_)
	var i Level
	err := row.Scan(&i.ID, &i.Map)
	return i, err
}

const createPlayer = `-- name: CreatePlayer :one
INSERT INTO players (
    level_id, start_x, start_y, curr_x, curr_y, hitpoints
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING id, hitpoints, start_x, start_y, curr_x, curr_y, level_id
`

type CreatePlayerParams struct {
	LevelID   int32
	StartX    int32
	StartY    int32
	CurrX     int32
	CurrY     int32
	Hitpoints int32
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (Player, error) {
	row := q.db.QueryRowContext(ctx, createPlayer,
		arg.LevelID,
		arg.StartX,
		arg.StartY,
		arg.CurrX,
		arg.CurrY,
		arg.Hitpoints,
	)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Hitpoints,
		&i.StartX,
		&i.StartY,
		&i.CurrX,
		&i.CurrY,
		&i.LevelID,
	)
	return i, err
}

const getLevel = `-- name: GetLevel :one
SELECT id, map FROM levels
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetLevel(ctx context.Context, id int32) (Level, error) {
	row := q.db.QueryRowContext(ctx, getLevel, id)
	var i Level
	err := row.Scan(&i.ID, &i.Map)
	return i, err
}

const getPlayerByLevelIDAndLock = `-- name: GetPlayerByLevelIDAndLock :one
SELECT id, hitpoints, start_x, start_y, curr_x, curr_y, level_id FROM players WHERE level_id = $1 FOR UPDATE
`

func (q *Queries) GetPlayerByLevelIDAndLock(ctx context.Context, levelID int32) (Player, error) {
	row := q.db.QueryRowContext(ctx, getPlayerByLevelIDAndLock, levelID)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Hitpoints,
		&i.StartX,
		&i.StartY,
		&i.CurrX,
		&i.CurrY,
		&i.LevelID,
	)
	return i, err
}

const updatePlayer = `-- name: UpdatePlayer :execrows
UPDATE players
SET hitpoints = $2, curr_x = $3, curr_y = $4
WHERE level_id = $1
`

type UpdatePlayerParams struct {
	LevelID   int32
	Hitpoints int32
	CurrX     int32
	CurrY     int32
}

func (q *Queries) UpdatePlayer(ctx context.Context, arg UpdatePlayerParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, updatePlayer,
		arg.LevelID,
		arg.Hitpoints,
		arg.CurrX,
		arg.CurrY,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
