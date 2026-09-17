package handler

import "shortly/internal/apperr"

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
