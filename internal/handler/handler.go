package handler

import "shortly/internal/store"

const SCHEME = "http"
const DOMAIN = "example.com"
const URL_PREFIX = SCHEME + "://" + DOMAIN + "/"

type Handler struct {
	store *store.Store
}

func New(s *store.Store) *Handler {
	return &Handler{s}
}
