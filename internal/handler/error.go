package handler

import (
	"encoding/json/v2"
	"errors"
	"log/slog"
	"net/http"
	"shortly/internal/apperr"
)

var (
	ErrMalformedBody = apperr.New("MALFORMED_BODY", "unable to parse the request")
	ErrInternal      = apperr.New("INTERNAL", "internal server error")
)

var statusFromCode = map[string]int{
	apperr.ErrAlreadyExists.Code: http.StatusConflict,
	apperr.ErrNotFound.Code:      http.StatusNotFound,
	ErrInvalidAlias.Code:         http.StatusUnprocessableEntity,
	ErrInvalidAliasLength.Code:   http.StatusUnprocessableEntity,
	ErrInvalidURLLength.Code:     http.StatusUnprocessableEntity,
	ErrMissingHost.Code:          http.StatusUnprocessableEntity,
	ErrInvalidHost.Code:          http.StatusUnprocessableEntity,
	ErrMalformedBody.Code:        http.StatusBadRequest,
}

//These functions should never be called on nil errors.
//TODO : Refactor for better handling of error type assertion.

func getStatus(err error) int {
	var ae *apperr.Error
	if errors.As(err, &ae) {
		if status, ok := statusFromCode[ae.Code]; ok {
			return status
		}
	}
	return http.StatusInternalServerError
}

func logError(err error) {
	var ae *apperr.Error
	if errors.As(err, &ae) {
		slog.Error(ae.Code + " " + ae.Message)
		return
	}
	slog.Error("UNDEFINED_ERROR " + err.Error())
}

func WritePageError(w http.ResponseWriter, e error) {
	//TODO: serve HTML error page built by astro
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func WriteAPIError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(getStatus(err))
	var ae *apperr.Error
	//Unhandled JSON marshal errors
	if errors.As(err, &ae) {
		json.MarshalWrite(w, err)
		return
	}
	json.MarshalWrite(w, ErrInternal)
	logError(err)
}
