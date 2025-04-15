package gopie_test

import (
	"gopie"
	"testing"
)

func TestTruncateString(t *testing.T) {
	shouldEqual(t, gopie.TruncateString("ABC👍DEFG", 4), "ABC👍...")
	shouldEqual(t, gopie.TruncateString("ABC👍DEFG", 8), "ABC👍DEFG")
}
