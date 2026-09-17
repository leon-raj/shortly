package handler

import (
	"net/http"
	"shortly/internal/apperr"
)

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {

	alias := r.PathValue("alias")
	if err := checkValidAlias(alias); err != nil {
		WritePageError(w, apperr.ErrNotFound)
		return
	}

	url, err := h.store.GetRedirectURL(alias)
	if err != nil {
		WritePageError(w, err)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}
