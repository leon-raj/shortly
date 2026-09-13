package handler

import (
	"errors"
	"net/http"
	"shortly/internal/store"
)

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {

	s := r.PathValue("shortCode")
	if !isValidShortCode(s) {
		http.Error(w, "invalid short code", http.StatusBadRequest)
		return
	}

	url, err := h.store.GetRedirectURL(s)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
