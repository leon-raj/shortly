package handler

import "shortly/internal/store"

type Handler struct {
	store *store.Store
}

const SCHEME = "http"
const DOMAIN = "example.com"
const URL_PREFIX = SCHEME + "://" + DOMAIN + "/"
