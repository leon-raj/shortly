package handler

import (
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
		if err == store.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
