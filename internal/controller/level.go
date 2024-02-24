package controller

import (
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

//--------------------------------------------------------------------------------
// LevelController
//--------------------------------------------------------------------------------

// LevelController handles HTTP requests for level management.
type LevelController struct {
	db     *gorm.DB
	logger *zap.Logger
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

// SubmitLevel handles HTTP requests to insert levels that can be played.
// It returns the unique ID of the level or an error.
func (lc *LevelController) SubmitLevel(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	cells := make([][]model.Cell, 0)

	if err := deserializePostRequest(w, r, &cells); err != nil {
		handleError(lc.logger, w, err)
		return
	}

	lc.logger.Debug("unmarshalled_level", zap.Any("cells", cells))

	lvl, err := model.NewLevel(cells)

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

func (lc *LevelController) createMap(lvl *model.Level) error {
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

//--------------------------------------------------------------------------------
// InsertLevelResponse
//--------------------------------------------------------------------------------

type InsertLevelResponse struct {
	Id      uint   `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}
