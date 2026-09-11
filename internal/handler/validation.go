package handler

const MAXLEN = 32

func isBase62(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9')) {
			return false
		}
	}
	return true
}

func isValidShortCode(s string) bool {
	return len(s) <= MAXLEN && isBase62(s)
}
