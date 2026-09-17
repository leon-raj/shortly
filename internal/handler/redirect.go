package handler

import (
	"net/http"
	"shortly/internal/apperr"
)

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {

	s := r.PathValue("shortCode")
	if err := checkValidShortCode(s); err != nil {
		WritePageError(w, apperr.ErrNotFound)
		return
	}

	url, err := h.store.GetRedirectURL(s)
	if err != nil {
		WritePageError(w, err)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
