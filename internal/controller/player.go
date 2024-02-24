package controller

import (
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

type PlayerController struct {
	db     *gorm.DB
	logger *zap.Logger
}

type MovePlayerResponse struct {
	Id      uint   `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
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

}
