package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/JakobDFrank/penn-roguelike/server/internal/analytics"
	"github.com/JakobDFrank/penn-roguelike/server/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/server/internal/database/cache"
	"go.uber.org/zap"
	"strconv"
	"time"
)

const (
	_expirationTime = time.Minute * 10
)

//--------------------------------------------------------------------------------
// sqlcLevelRepository
//--------------------------------------------------------------------------------

type sqlcLevelRepository struct {
	db      *sql.DB
	queries *Queries
	cache   cache.Cacher
	obs     analytics.Collector
	logger  *zap.Logger
}

// NewSqlcLevelRepository returns a SQLC implementation of LevelRepository.
func NewSqlcLevelRepository(db *sql.DB, logger *zap.Logger, obs analytics.Collector, cache cache.Cacher) (LevelRepository, error) {

	if db == nil {
		return nil, &apperr.NilArgumentError{Message: "db"}
	}

	if obs == nil {
		return nil, &apperr.NilArgumentError{Message: "obs"}
	}

	if cache == nil {
		return nil, &apperr.NilArgumentError{Message: "cache"}
	}

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	queries := New(db)
	lr := sqlcLevelRepository{
		db:      db,
		queries: queries,
		obs:     obs,
		cache:   cache,
		logger:  logger,
	}

	return &lr, nil
}

func (lr *sqlcLevelRepository) Begin() (*sql.Tx, error) {
	return lr.db.Begin()
}

func (lr *sqlcLevelRepository) CreateLevelWithTx(tx *sql.Tx, cells [][]Cell) (Level, error) {

	defer analytics.MeasureDuration(lr.obs, "CreateLevelWithTx")()

	trans := lr.queries.WithTx(tx)
	jsonText, err := json.Marshal(cells)
	if err != nil {
		return Level{}, err
	}

	return trans.CreateLevel(context.Background(), jsonText)
}

func (lr *sqlcLevelRepository) GetFirst(id int32) (Level, error) {

	defer analytics.MeasureDuration(lr.obs, "GetFirst")()

	ctx := context.Background()
	key := "level:" + strconv.FormatInt(int64(id), 32)

	// Levels are immutable, so we don't need to worry about locking here
	lvl, err := cache.GetStruct[Level](lr.cache, ctx, key)

	if err == nil {
		lr.logger.Info("cache_hit")
		return lvl, nil
	}

	lr.logger.Info("cache_miss")

	// cache miss or doesn't exist
	lev, err := lr.queries.GetLevel(ctx, id)

	if err != nil {
		return lev, err
	}

	lr.logger.Info("set_cache", zap.String("key", key))
	if err := cache.SetStruct(lr.cache, ctx, key, lev, _expirationTime); err != nil {
		lr.logger.Error("set_cache", zap.Error(err))
		// dont return err
	}

	return lev, nil
}

var _ LevelRepository = (*sqlcLevelRepository)(nil)

//--------------------------------------------------------------------------------
// sqlcPlayerRepository
//--------------------------------------------------------------------------------

type sqlcPlayerRepository struct {
	db      *sql.DB
	queries *Queries
	obs     analytics.Collector
}

// NewSqlcPlayerRepository returns a SQLC implementation of PlayerRepository.
func NewSqlcPlayerRepository(db *sql.DB, obs analytics.Collector) (PlayerRepository, error) {

	if db == nil {
		return nil, &apperr.NilArgumentError{Message: "db"}
	}

	queries := New(db)
	pr := sqlcPlayerRepository{
		db:      db,
		queries: queries,
		obs:     obs,
	}

	return &pr, nil
}

func (pr *sqlcPlayerRepository) Begin() (*sql.Tx, error) {
	return pr.db.Begin()
}

func (pr *sqlcPlayerRepository) CreatePlayerWithTx(tx *sql.Tx, params CreatePlayerParams) (Player, error) {

	defer analytics.MeasureDuration(pr.obs, "CreatePlayerWithTx")()

	trans := pr.queries.WithTx(tx)
	return trans.CreatePlayer(context.Background(), params)
}

func (pr *sqlcPlayerRepository) GetPlayerByLevelIDAndLockWithTx(tx *sql.Tx, levelId int32) (Player, error) {
	defer analytics.MeasureDuration(pr.obs, "GetPlayerByLevelIDAndLockWithTx")()

	trans := pr.queries.WithTx(tx)
	return trans.GetPlayerByLevelIDAndLock(context.Background(), levelId)
}

func (pr *sqlcPlayerRepository) UpdatePlayer(tx *sql.Tx, params UpdatePlayerParams) error {

	defer analytics.MeasureDuration(pr.obs, "UpdatePlayer")()

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

var _ PlayerRepository = (*sqlcPlayerRepository)(nil)
