package handler

import (
	"encoding/json/v2"
	"errors"
	"log/slog"
	"net/http"
	"shortly/internal/store"
)

type Request struct {
	ShortCode string `json:"shortCode"`
	URL       string `json:"url"`
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {

	var req Request
	err := json.UnmarshalRead(r.Body, &req)
	if err != nil {
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}
	if !isValidShortCode(req.ShortCode) {
		http.Error(w, "invalid short code", http.StatusBadRequest)
		return
	}
	url, err := sanitizeURL(req.URL)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidScheme):
			http.Error(w, "scheme must be http or https", http.StatusBadRequest)
		case errors.Is(err, ErrInvalidLength):
			http.Error(w, "URL must be non-empty with length less than or equal to 2048 characters", http.StatusBadRequest)
		case errors.Is(err, ErrMissingHost):
			http.Error(w, "URL must contain a host", http.StatusBadRequest)
		case errors.Is(err, ErrInvalidHost):
			http.Error(w, "hostname not allowed", http.StatusBadRequest)
		default:
			http.Error(w, "invalid URL", http.StatusBadRequest)
		}
		return
	}
	err = h.store.PutRedirectURL(req.ShortCode, url)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrAlreadyExists):
			http.Error(w, "short code already in use", http.StatusConflict)
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", URL_PREFIX+req.ShortCode)
	w.WriteHeader(http.StatusCreated)
	//while unlikely, if the marshaling step fail, the client will receive StatusCreated with an improper body.
	err = json.MarshalWrite(w, Request{req.ShortCode, url})
	if err != nil {
		slog.Error("failed to marshal")
	}
}
