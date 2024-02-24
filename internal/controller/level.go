package controller

import (
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/model/level"
	"github.com/JakobDFrank/penn-roguelike/internal/model/player"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type LevelController struct {
	db     *gorm.DB
	logger *zap.Logger
}

type InsertLevelResponse struct {
	Id      uint   `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewLevelController(logger *zap.Logger, db *gorm.DB) (*LevelController, error) {
	if db == nil {
		return nil, &apperr.NilArgumentError{Message: "db"}
	}

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	lc := &LevelController{
		db:     db,
		logger: logger,
	}

	return lc, nil
}

func (lc *LevelController) SubmitLevel(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(lc.logger, w, err)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			lc.logger.Error("close_body", zap.Error(err))
		}
	}()

	cells := make([][]level.Cell, 0)
	if err := json.Unmarshal(body, &cells); err != nil {
		handleError(lc.logger, w, err)
		return
	}

	lc.logger.Debug("unmarshalled_level", zap.Any("cells", cells))

	lvl, err := level.NewLevel(cells)

	if err != nil {
		handleError(lc.logger, w, err)
		return
	}

	if err := lc.createMap(lvl); err != nil {
		handleError(lc.logger, w, err)
		return
	}

	resp := InsertLevelResponse{
		Id:      lvl.ID,
		Message: "",
		Status:  http.StatusOK,
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

func (lc *LevelController) createMap(lvl *level.Level) error {
	tx := lc.db.Begin()

	lvlRes := tx.Create(lvl)

	if lvlRes.Error != nil {
		tx.Rollback()
		return lvlRes.Error
	}

	playr := player.NewPlayer(lvl.ID, lvl.RowStart, lvl.ColStart)

	playerRes := tx.Create(playr)

	if playerRes.Error != nil {
		tx.Rollback()
		return playerRes.Error
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	lc.logger.Debug("submit_level", zap.Int64("lvl_rows_affected", lvlRes.RowsAffected), zap.Int64("player_rows_affected", playerRes.RowsAffected))

	fmt.Println(lvl.Cells.String())

	return nil
}

func handleError(logger *zap.Logger, w http.ResponseWriter, err error) {

	logger.Error("handling_error", zap.Error(err))

	resp := InsertLevelResponse{
		Id:      0,
		Message: err.Error(),
		Status:  http.StatusBadRequest,
	}

	jsonData, err := json.Marshal(resp)

	if err != nil {
		logger.Error("marshal_error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(jsonData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
