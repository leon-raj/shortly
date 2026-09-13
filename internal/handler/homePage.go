package handler

import "net/http"

func (h *Handler) homePage(w http.ResponseWriter, _ http.Request) {
	http.Error(w, "yet to implement", http.StatusNotImplemented)
}
