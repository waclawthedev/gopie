package gopie

import (
	"unicode"
	"unicode/utf8"
)

// TruncateString returns string truncated by maxRunes runes with "..." at the end.
func TruncateString(s string, maxRunes int) string {
	if utf8.RuneCountInString(s) <= maxRunes {
		return s
	}

	return string([]rune(s)[:maxRunes]) + "..."
}

// WithoutInvisibleChars returns string without invisible characters.
// IsPrint&IsGraphic avoided to bypass bug described here: https://github.com/golang/go/issues/73673
func WithoutInvisibleChars(s string) string {
	var output []rune

	for _, char := range s {
		if IsVisible(char) {
			output = append(output, char)
		}
	}

	return string(output)
}

// IsVisible returns true if rune is visible. Control characters defined as invisible.
// IsPrint&IsGraphic avoided to bypass bug described here: https://github.com/golang/go/issues/73673
func IsVisible(char rune) bool {
	return !unicode.IsControl(char) && (unicode.IsNumber(char) ||
		unicode.IsDigit(char) ||
		unicode.IsLetter(char) ||
		unicode.IsSpace(char) ||
		unicode.IsPunct(char) ||
		unicode.IsSymbol(char))
}
