package gopie_test

import (
	"testing"

	"github.com/waclawthedev/gopie"
)

func TestTruncateString(t *testing.T) {
	shouldEqual(t, gopie.TruncateString("ABC👍DEFG", 4), "ABC👍...")
	shouldEqual(t, gopie.TruncateString("ABC👍DEFG", 8), "ABC👍DEFG")
}

func TestWithoutInvisibleChars(t *testing.T) {
	shouldEqual(t, gopie.WithoutInvisibleChars("ABC👍DEFG"), "ABC👍DEFG")
	shouldEqual(t, gopie.WithoutInvisibleChars("§1234567890-=qwertyuiop[]asdfghjkl;'\\`zxcvbnm,./"),
		"§1234567890-=qwertyuiop[]asdfghjkl;'\\`zxcvbnm,./")
	shouldEqual(t, gopie.WithoutInvisibleChars("1234567890-=йцукенгшщзхїфівапролджєʼґячсмитьбю."),
		"1234567890-=йцукенгшщзхїфівапролджєʼґячсмитьбю.")
	shouldEqual(t, gopie.WithoutInvisibleChars("ą, ć, ę, ł, ń, ó, ś, ź, ż"), "ą, ć, ę, ł, ń, ó, ś, ź, ż")

	shouldEqual(t, gopie.WithoutInvisibleChars("1\n"), "1")
}

func TestIsVisible(t *testing.T) {
	shouldEqual(t, gopie.IsVisible('1'), true)
	shouldEqual(t, gopie.IsVisible('a'), true)
	shouldEqual(t, gopie.IsVisible('Ї'), true)
	shouldEqual(t, gopie.IsVisible('Ї'), true)
	shouldEqual(t, gopie.IsVisible('語'), true)
	shouldEqual(t, gopie.IsVisible('"'), true)
	shouldEqual(t, gopie.IsVisible('&'), true)
	shouldEqual(t, gopie.IsVisible('@'), true)

	shouldEqual(t, gopie.IsVisible('\n'), false)
	shouldEqual(t, gopie.IsVisible('\t'), false)
	shouldEqual(t, gopie.IsVisible('󠀮'), false) // U+E002E TAG FULL STOP
	shouldEqual(t, gopie.IsVisible('󠇯'), false) // U+E01EF VARIATION SELECTOR-256
	shouldEqual(t, gopie.IsVisible('󠅗'), false) // U+E0157 VARIATION SELECTOR-104
	shouldEqual(t, gopie.IsVisible('󠁾'), false) // U+E007E TAG TILDE
}
