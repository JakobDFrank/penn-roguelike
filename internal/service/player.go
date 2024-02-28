package service

import (
	"encoding/json"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/database/model"
	"go.uber.org/zap"
	"time"
)

//--------------------------------------------------------------------------------
// PlayerService
//--------------------------------------------------------------------------------

const (
	_startingHitpoints = 4
)

// PlayerService handles player management.
type PlayerService struct {
	levelRepo  model.LevelRepository
	playerRepo model.PlayerRepository
	logger     *zap.Logger
}

// NewPlayerService creates a new instance of PlayerService.
func NewPlayerService(logger *zap.Logger, levelRepo model.LevelRepository, playerRepo model.PlayerRepository) (*PlayerService, error) {

	if levelRepo == nil {
		return nil, &apperr.NilArgumentError{Message: "levelRepo"}
	}

	if playerRepo == nil {
		return nil, &apperr.NilArgumentError{Message: "playerRepo"}
	}

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	pc := &PlayerService{
		levelRepo:  levelRepo,
		playerRepo: playerRepo,
		logger:     logger,
	}

	return pc, nil
}

// MovePlayer will attempt to move a player on a map in a given direction.
// It returns the new game state or an error.
func (pc *PlayerService) MovePlayer(id int32, dir model.Direction) (*model.Level, error) {

	lvl, err := pc.movePlayer(id, dir)

	if err != nil {
		return nil, err
	}

	return lvl, nil
}

func (pc *PlayerService) movePlayer(id int32, dir model.Direction) (*model.Level, error) {

	pc.logger.Debug("move_player", zap.String("dir", dir.String()))

	lvl, err := pc.levelRepo.GetFirst(id)

	if err != nil {
		return nil, err
	}

	cells := make([][]model.Cell, 0)
	if err := json.Unmarshal(lvl.Map, &cells); err != nil {
		return nil, err
	}

	tx, err := pc.playerRepo.Begin()

	if err != nil {
		return nil, err
	}

	playr, err := pc.playerRepo.GetPlayerByLevelIDAndLockWithTx(tx, lvl.ID)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	start := time.Now()
	pc.logger.Debug("start_lock", zap.Time("start_time", start))

	moved, err := pc.tryMove(cells, &playr, dir)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if moved {
		if err := pc.playerRepo.UpdatePlayer(tx, model.UpdatePlayerParams{
			LevelID:   playr.LevelID,
			Hitpoints: playr.Hitpoints,
			CurrX:     playr.CurrX,
			CurrY:     playr.CurrY,
		}); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	end := time.Now()
	pc.logger.Debug("ending_lock", zap.Time("end_time", end), zap.Duration("elapsed", end.Sub(start)))

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	data, err := SerializeCellsWithPlayer(cells, &playr, pc.logger)

	if err != nil {
		return nil, err
	}

	if err := lvl.Map.UnmarshalJSON(data); err != nil {
		return nil, err
	}

	return &lvl, nil
}

func (pc *PlayerService) tryMove(lvl [][]model.Cell, p *model.Player, dir model.Direction) (bool, error) {
	switch dir {
	case model.Left:
		if p.CurrX > 0 {
			moved := pc.handlePlayerMoveAttempt(lvl, p, p.CurrY, p.CurrX-1)
			return moved, nil
		}
	case model.Right:
		if int(p.CurrX) < len(lvl[0])-1 {
			moved := pc.handlePlayerMoveAttempt(lvl, p, p.CurrY, p.CurrX+1)
			return moved, nil
		}
	case model.Up:
		if p.CurrY > 0 {
			moved := pc.handlePlayerMoveAttempt(lvl, p, p.CurrY-1, p.CurrX)
			return moved, nil
		}
	case model.Down:
		if int(p.CurrY) < len(lvl)-1 {
			moved := pc.handlePlayerMoveAttempt(lvl, p, p.CurrY+1, p.CurrX)
			return moved, nil
		}
	default:
		return false, &apperr.UnimplementedError{}
	}

	pc.logger.Info("player_tried_moving_out_of_bounds")
	return false, nil
}

func (pc *PlayerService) handlePlayerMoveAttempt(lvl [][]model.Cell, p *model.Player, row, col int32) bool {
	switch lvl[row][col] {
	case model.CellWall:
		pc.logger.Info("player_blocked_by_wall")
		return false
	case model.CellPit:
		p.Hitpoints -= 1
		pc.logger.Info("player_fell_into_pit", zap.Int32("hp", p.Hitpoints))
	case model.CellArrow:
		p.Hitpoints -= 2
		pc.logger.Info("player_hit_by_arrows", zap.Int32("hp", p.Hitpoints))
	case model.CellOpen, model.CellPlayer:
		pc.logger.Info("player_moved", zap.Int32("row", row), zap.Int32("col", col))
	}

	if p.Hitpoints > 0 {
		p.CurrY = row
		p.CurrX = col
	} else {
		p.CurrY = p.StartY
		p.CurrX = p.StartX
		p.Hitpoints = _startingHitpoints

		pc.logger.Info("player_died", zap.Int32("x", p.CurrX), zap.Int32("y", p.CurrY))
	}

	return true
}
