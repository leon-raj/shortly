package handler

import (
	"encoding/json/v2"
	"log/slog"
	"net/http"
)

type Request struct {
	URL string `json:"url"`
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {

	shortCode := r.PathValue("shortCode")
	if err := checkValidShortCode(shortCode); err != nil {
		WriteAPIError(w, err)
		return
	}
	var req Request
	if err := json.UnmarshalRead(r.Body, &req); err != nil {
		WriteAPIError(w, ErrMalformedBody)
		return
	}
	url, err := sanitizeURL(req.URL)
	if err != nil {
		WriteAPIError(w, err)
		return
	}
	err = h.store.PutRedirectURL(shortCode, url)
	if err != nil {
		WriteAPIError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", URL_PREFIX+shortCode)
	w.WriteHeader(http.StatusCreated)
	//while unlikely, if the marshaling step fail, the client will receive StatusCreated with an improper body.
	err = json.MarshalWrite(w, Request{url})
	if err != nil {
		slog.Error("failed to marshal")
	}
}
