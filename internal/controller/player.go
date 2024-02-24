package controller

import (
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/model/level"
	"github.com/JakobDFrank/penn-roguelike/internal/model/player"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"net/http"
)

type PlayerController struct {
	db     *gorm.DB
	logger *zap.Logger
}

type MovePlayerResponse struct {
	Id     uint   `json:"id"`
	Level  string `json:"level"`
	Status int    `json:"status"`
}

func NewPlayerController(logger *zap.Logger, db *gorm.DB) (*PlayerController, error) {
	if db == nil {
		return nil, &apperr.NilArgumentError{Message: "db"}
	}

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	pc := &PlayerController{
		db:     db,
		logger: logger,
	}

	return pc, nil
}

func (pc *PlayerController) MovePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(pc.logger, w, err)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			pc.logger.Error("close_body", zap.Error(err))
		}
	}()

	moveRequest := MovePlayerRequest{}
	if err := json.Unmarshal(body, &moveRequest); err != nil {
		handleError(pc.logger, w, err)
		return
	}

	pc.logger.Debug("unmarshal", zap.Any("move_request", moveRequest))

	lvl, err := pc.movePlayer(moveRequest)

	if err != nil {
		handleError(pc.logger, w, err)
	}

	cellJson, err := json.Marshal(lvl.Cells)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := MovePlayerResponse{
		Id:     moveRequest.ID,
		Level:  string(cellJson),
		Status: http.StatusOK,
	}

	jsonData, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(jsonData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (pc *PlayerController) movePlayer(req MovePlayerRequest) (*level.Level, error) {

	dir, err := player.NewDirection(req.Direction)

	if err != nil {
		return nil, err
	}

	pc.logger.Debug("move_player", zap.String("dir", dir.String()))

	var lvl level.Level

	if err := pc.db.First(&lvl, req.ID).Error; err != nil {
		return nil, err
	}

	tx := pc.db.Begin()

	var playr player.Player

	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("level_id = ?", req.ID).First(&playr).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	oldRowIdx := playr.RowIdx
	oldColIdx := playr.ColIdx

	lvl.Cells[lvl.RowStart][lvl.ColStart] = level.Open
	lvl.Cells[oldRowIdx][oldColIdx] = level.Player
	fmt.Printf("before: \n%s\n", lvl.Cells.String())

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

	lvl.Cells[oldRowIdx][oldColIdx] = level.Open
	lvl.Cells[playr.RowIdx][playr.ColIdx] = level.Player

	fmt.Printf("after: \n%s\n", lvl.Cells.String())

	return &lvl, nil
}

func (pc *PlayerController) tryMove(lvl *level.Level, p *player.Player, dir player.Direction) (bool, error) {
	switch dir {
	case player.Left:
		if p.ColIdx > 0 {
			moved := pc.handlePlayerMoveAttempt(lvl, p, p.RowIdx, p.ColIdx-1)
			return moved, nil
		}
	case player.Right:
		if p.ColIdx < len(lvl.Cells[0])-1 {
			moved := pc.handlePlayerMoveAttempt(lvl, p, p.RowIdx, p.ColIdx+1)
			return moved, nil
		}
	case player.Up:
		if p.RowIdx > 0 {
			moved := pc.handlePlayerMoveAttempt(lvl, p, p.RowIdx-1, p.ColIdx)
			return moved, nil
		}
	case player.Down:
		if p.RowIdx < len(lvl.Cells)-1 {
			moved := pc.handlePlayerMoveAttempt(lvl, p, p.RowIdx+1, p.ColIdx)
			return moved, nil
		}
	default:
		return false, &apperr.UnimplementedError{}
	}

	pc.logger.Info("player_tried_moving_out_of_bounds")
	return false, nil
}

func (pc *PlayerController) handlePlayerMoveAttempt(lvl *level.Level, p *player.Player, row, col int) bool {
	switch lvl.Cells[row][col] {
	case level.Wall:
		pc.logger.Info("player_blocked_by_wall")
		return false
	case level.Pit:
		p.Hitpoints -= 1
		pc.logger.Info("player_fell_into_pit", zap.Int("hp", p.Hitpoints))
	case level.Arrow:
		p.Hitpoints -= 2
		pc.logger.Info("player_hit_by_arrows", zap.Int("hp", p.Hitpoints))
	case level.Open, level.Player:
		pc.logger.Info("player_moved", zap.Int("row", row), zap.Int("col", col))
	}

	if p.Hitpoints > 0 {
		p.RowIdx = row
		p.ColIdx = col
	} else {
		p.RowIdx = lvl.RowStart
		p.ColIdx = lvl.ColStart
		p.ResetHitpoints()

		pc.logger.Info("player_died", zap.Int("row", p.RowIdx), zap.Int("col", p.ColIdx))
	}

	return true
}

type MovePlayerRequest struct {
	ID        uint `json:"id"`
	Direction int  `json:"direction"`
}
