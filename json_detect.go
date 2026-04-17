package wire

// looksLikeJSON returns true if s looks like it could be a JSON value.
// It skips leading whitespace and checks whether the first non-whitespace
// character is one that can start a valid JSON value:
//
//	{ [ " - t f n 0-9
//
// This is a heuristic — it does not validate the entire string.
func looksLikeJSON(s string) bool {
	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
			continue
		}
		switch r {
		case '{', '[', '"', '-', 't', 'f', 'n':
			return true
		}
		if r >= '0' && r <= '9' {
			return true
		}
		return false
	}
	return false
}
