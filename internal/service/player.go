package service

import (
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

//--------------------------------------------------------------------------------
// PlayerService
//--------------------------------------------------------------------------------

// PlayerService handles player management.
type PlayerService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPlayerController(logger *zap.Logger, db *gorm.DB) (*PlayerService, error) {
	if db == nil {
		return nil, &apperr.NilArgumentError{Message: "db"}
	}

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	pc := &PlayerService{
		db:     db,
		logger: logger,
	}

	return pc, nil
}

// MovePlayer will attempt to move a player on a map in a given direction.
// It returns the new game state or an error.
func (pc *PlayerService) MovePlayer(id uint, dir model.Direction) (string, error) {

	lvl, err := pc.movePlayer(id, dir)

	if err != nil {
		return "", err
	}

	cellJson, err := json.Marshal(lvl.Map)

	if err != nil {
		return "", err
	}

	return string(cellJson), err
}

func (pc *PlayerService) movePlayer(id uint, dir model.Direction) (*model.Level, error) {

	pc.logger.Debug("move_player", zap.String("dir", dir.String()))

	var lvl model.Level

	if err := pc.db.First(&lvl, id).Error; err != nil {
		return nil, err
	}

	tx := pc.db.Begin()

	start := time.Now()
	pc.logger.Debug("start_transaction", zap.Time("start_time", start))

	var playr model.Player

	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("level_id = ?", id).First(&playr).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	oldRowIdx := playr.RowIdx
	oldColIdx := playr.ColIdx

	lvl.Map[lvl.RowStartIdx][lvl.ColStartIdx] = model.CellOpen
	lvl.Map[oldRowIdx][oldColIdx] = model.CellPlayer
	fmt.Printf("before: \n%s\n", lvl.Map.String())

	moved, err := pc.tryMove(&lvl, &playr, dir)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if moved {
		if err := tx.Save(&playr).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	end := time.Now()
	pc.logger.Debug("end_transaction", zap.Time("end_time", end), zap.Duration("elapsed", end.Sub(start)))

	lvl.Map[oldRowIdx][oldColIdx] = model.CellOpen
	lvl.Map[playr.RowIdx][playr.ColIdx] = model.CellPlayer

	fmt.Printf("after: \n%s\n", lvl.Map.String())

	return &lvl, nil
}

func (pc *PlayerService) tryMove(lvl *model.Level, p *model.Player, dir model.Direction) (bool, error) {
	switch dir {
	case model.Left:
		if p.ColIdx > 0 {
			moved := pc.handlePlayerMoveAttempt(lvl, p, p.RowIdx, p.ColIdx-1)
			return moved, nil
		}
	case model.Right:
		if p.ColIdx < len(lvl.Map[0])-1 {
			moved := pc.handlePlayerMoveAttempt(lvl, p, p.RowIdx, p.ColIdx+1)
			return moved, nil
		}
	case model.Up:
		if p.RowIdx > 0 {
			moved := pc.handlePlayerMoveAttempt(lvl, p, p.RowIdx-1, p.ColIdx)
			return moved, nil
		}
	case model.Down:
		if p.RowIdx < len(lvl.Map)-1 {
			moved := pc.handlePlayerMoveAttempt(lvl, p, p.RowIdx+1, p.ColIdx)
			return moved, nil
		}
	default:
		return false, &apperr.UnimplementedError{}
	}

	pc.logger.Info("player_tried_moving_out_of_bounds")
	return false, nil
}

func (pc *PlayerService) handlePlayerMoveAttempt(lvl *model.Level, p *model.Player, row, col int) bool {
	switch lvl.Map[row][col] {
	case model.CellWall:
		pc.logger.Info("player_blocked_by_wall")
		return false
	case model.CellPit:
		p.Hitpoints -= 1
		pc.logger.Info("player_fell_into_pit", zap.Int("hp", p.Hitpoints))
	case model.CellArrow:
		p.Hitpoints -= 2
		pc.logger.Info("player_hit_by_arrows", zap.Int("hp", p.Hitpoints))
	case model.CellOpen, model.CellPlayer:
		pc.logger.Info("player_moved", zap.Int("row", row), zap.Int("col", col))
	}

	if p.Hitpoints > 0 {
		p.RowIdx = row
		p.ColIdx = col
	} else {
		p.RowIdx = lvl.RowStartIdx
		p.ColIdx = lvl.ColStartIdx
		p.ResetHitpoints()

		pc.logger.Info("player_died", zap.Int("row", p.RowIdx), zap.Int("col", p.ColIdx))
	}

	return true
}
