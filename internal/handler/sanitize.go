package handler

import (
	"errors"
	"net/url"
	"strings"
)

var (
	ErrInvalidLength = errors.New("invalid URL length")
	ErrInvalidScheme = errors.New("invalid scheme")
	ErrMissingHost   = errors.New("missing host")
	ErrInvalidHost   = errors.New("invalid host")
)

func sanitizeURL(raw string) (string, error) {
	if len(raw) == 0 || len(raw) > 2048 {
		return "", ErrInvalidLength
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
