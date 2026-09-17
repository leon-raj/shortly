package handler

import (
	"net/url"
	"shortly/internal/apperr"
	"strings"
)

const MAXLEN = 32

var (
	ErrInvalidAlias       = apperr.New("INVALID_ALIAS", "the shortcode contains invalid characters")
	ErrInvalidAliasLength = apperr.New("INVALID_ALIAS_LENGTH", "the shortcode should be less than or equal to 32 characters and non-empty")
)

func isBase62(s string) bool {
	for i := 0; i < len(s); i++ {
		if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9')) {
			return false
		}
	}
	return true
}

func checkValidAlias(alias string) error {
	if !(len(alias) > 0 && len(alias) <= MAXLEN) {
		return ErrInvalidAliasLength
	}
	if !isBase62(alias) {
		return ErrInvalidAlias
	}
	return nil
}

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
