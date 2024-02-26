package service

import (
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//--------------------------------------------------------------------------------
// LevelService
//--------------------------------------------------------------------------------

// LevelService handles level management.
type LevelService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewLevelController(logger *zap.Logger, db *gorm.DB) (*LevelService, error) {
	if db == nil {
		return nil, &apperr.NilArgumentError{Message: "db"}
	}

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	lc := &LevelService{
		db:     db,
		logger: logger,
	}

	return lc, nil
}

// SubmitLevel inserts levels that can be played.
// It returns the unique ID of the level or an error.
func (lc *LevelService) SubmitLevel(cells [][]model.Cell) (uint, error) {

	lc.logger.Debug("unmarshalled_level", zap.Any("cells", cells))

	lvl, err := model.NewLevel(cells)

	if err != nil {
		return 0, err
	}

	if err := lc.createMap(lvl); err != nil {
		return 0, err
	}

	return lvl.ID, nil
}

func (lc *LevelService) createMap(lvl *model.Level) error {
	tx := lc.db.Begin()

	lvlRes := tx.Create(lvl)

	if lvlRes.Error != nil {
		tx.Rollback()
		return lvlRes.Error
	}

	playr := model.NewPlayer(lvl.ID, lvl.RowStartIdx, lvl.ColStartIdx)

	playerRes := tx.Create(playr)

	if playerRes.Error != nil {
		tx.Rollback()
		return playerRes.Error
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	lc.logger.Debug("submit_level", zap.Int64("lvl_rows_affected", lvlRes.RowsAffected), zap.Int64("player_rows_affected", playerRes.RowsAffected))

	fmt.Println(lvl.Map.String())

	return nil
}
