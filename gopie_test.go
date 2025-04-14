package gopie_test

import "testing"

func shouldEqual[T comparable](t *testing.T, got, expected T) {
	t.Helper()

	if got != expected {
		t.Errorf("got (%v) != expected (%v)", got, expected)
	}
}
