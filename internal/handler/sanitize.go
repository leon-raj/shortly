package handler

import (
	"net/url"
	"shortly/internal/apperr"
	"strings"
)

var (
	ErrInvalidURLLength = apperr.New("INVALID_URL_LENGTH", "URL must be non-empty with length less than or equal to 2048 characters")
	ErrInvalidScheme    = apperr.New("INVALID_SCHEME", "scheme must be http or https")
	ErrMissingHost      = apperr.New("MISSING_HOST", "URL must contain a host")
	ErrInvalidHost      = apperr.New("INVALID_HOST", "hostname not allowed")
)

// Need to disallow reserved URLs like DOMAIN/shorten
func sanitizeURL(raw string) (string, error) {
	if len(raw) == 0 || len(raw) > 2048 {
		return "", ErrInvalidURLLength
	}
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}
	if !(u.Scheme == "http" || u.Scheme == "https") {
		return "", ErrInvalidScheme
	}
	if u.Hostname() == "" {
		return "", ErrMissingHost
	}
	if strings.ToLower(u.Hostname()) == DOMAIN {
		return "", ErrInvalidHost
	}
	return u.String(), nil
}
