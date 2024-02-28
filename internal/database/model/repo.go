package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
)

type LevelRepository interface {
	Begin() (*sql.Tx, error)
	CreateLevelWithTx(tx *sql.Tx, cells [][]Cell) (Level, error)
	GetFirst(id int32) (Level, error)
}

type PlayerRepository interface {
	Begin() (*sql.Tx, error)
	CreatePlayerWithTx(tx *sql.Tx, levelId CreatePlayerParams) (Player, error)
	GetPlayerByLevelIDAndLockWithTx(tx *sql.Tx, levelId int32) (Player, error)
	UpdatePlayer(tx *sql.Tx, params UpdatePlayerParams) error
}

//--------------------------------------------------------------------------------
// levelRepository
//--------------------------------------------------------------------------------

type levelRepository struct {
	db      *sql.DB
	queries *Queries
}

func NewLevelRepository(db *sql.DB) (LevelRepository, error) {

	if db == nil {
		return nil, &apperr.NilArgumentError{Message: "db"}
	}

	queries := New(db)
	lr := levelRepository{
		db:      db,
		queries: queries,
	}

	return &lr, nil
}

func (lr *levelRepository) Begin() (*sql.Tx, error) {
	return lr.db.Begin()
}

func (lr *levelRepository) CreateLevelWithTx(tx *sql.Tx, cells [][]Cell) (Level, error) {
	trans := lr.queries.WithTx(tx)
	jsonText, err := json.Marshal(cells)
	if err != nil {
		return Level{}, err
	}

	return trans.CreateLevel(context.Background(), jsonText)
}

func (lr *levelRepository) GetFirst(id int32) (Level, error) {
	return lr.queries.GetLevel(context.Background(), id)
}

var _ LevelRepository = (*levelRepository)(nil)

//--------------------------------------------------------------------------------
// playerRepository
//--------------------------------------------------------------------------------

type playerRepository struct {
	db      *sql.DB
	queries *Queries
}

func NewPlayerRepository(db *sql.DB) (PlayerRepository, error) {

	if db == nil {
		return nil, &apperr.NilArgumentError{Message: "db"}
	}

	queries := New(db)
	pr := playerRepository{
		db:      db,
		queries: queries,
	}

	return &pr, nil
}

func (pr *playerRepository) Begin() (*sql.Tx, error) {
	return pr.db.Begin()
}

func (pr *playerRepository) CreatePlayerWithTx(tx *sql.Tx, params CreatePlayerParams) (Player, error) {
	trans := pr.queries.WithTx(tx)
	return trans.CreatePlayer(context.Background(), params)
}

func (pr *playerRepository) GetPlayerByLevelIDAndLockWithTx(tx *sql.Tx, levelId int32) (Player, error) {
	trans := pr.queries.WithTx(tx)
	return trans.GetPlayerByLevelIDAndLock(context.Background(), levelId)
}

func (pr *playerRepository) UpdatePlayer(tx *sql.Tx, params UpdatePlayerParams) error {
	trans := pr.queries.WithTx(tx)
	rows, err := trans.UpdatePlayer(context.Background(), params)

	if err != nil {
		return err
	}

	if rows == 0 {
		return &apperr.InvalidOperationError{Message: "player not found"}
	}

	return nil
}

var _ PlayerRepository = (*playerRepository)(nil)
