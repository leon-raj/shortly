package handler

import "net/http"

func (h *Handler) HomePage(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "yet to implement", http.StatusNotImplemented)
}
