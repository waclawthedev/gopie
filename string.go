package gopie

import "unicode/utf8"

// TruncateString returns string truncated by maxRunes runes with "..." at the end.
func TruncateString(s string, maxRunes int) string {
	if utf8.RuneCountInString(s) <= maxRunes {
		return s
	}

	return string([]rune(s)[:maxRunes]) + "..."
}
