package handler

import "shortly/internal/apperr"

const MAXLEN = 32

var (
	ErrInvalidShortCode       = apperr.New("INVALID_SHORTCODE", "the shortcode contains invalid characters")
	ErrInvalidShortCodeLength = apperr.New("INVALID_SHORT_CODE_LENGTH", "the shortcode should be less than or equal to 32 characters and non-empty")
)

func isBase62(s string) bool {
	for i := 0; i < len(s); i++ {
		if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9')) {
			return false
		}
	}
	return true
}

func checkValidShortCode(s string) error {
	if !(len(s) > 0 && len(s) <= MAXLEN) {
		return ErrInvalidShortCodeLength
	}
	if !isBase62(s) {
		return ErrInvalidShortCode
	}
	return nil
}
