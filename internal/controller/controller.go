// Package controller contains objects that handle HTTP requests to manage game actions.
package controller

import (
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"go.uber.org/zap"
	"net/http"
)

const (
	Kibibyte = 1024
	Mebibyte = Kibibyte * Kibibyte
)

func deserializePostRequest(w http.ResponseWriter, r *http.Request, value any) error {
	if r.Method != http.MethodPost {
		return &apperr.InvalidArgumentError{Message: fmt.Sprintf("invalid_method: %s", r.Method)}
	}

	// debug only? reflection a bit slow
	//if reflect.ValueOf(value).Kind() != reflect.Ptr {
	//	return &apperr.InvalidArgumentError{Message: "value must be a pointer"}
	//}

	// limit request size
	r.Body = http.MaxBytesReader(w, r.Body, Mebibyte)

	err := json.NewDecoder(r.Body).Decode(value)
	if err != nil {
		return err
	}

	return r.Body.Close()
}

func handleError(logger *zap.Logger, w http.ResponseWriter, err error) {

	logger.Error("handling_error", zap.Error(err))

	resp := ErrorResponse{
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

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
