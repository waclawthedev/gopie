package gopie_test

import (
	"gopie"
	"testing"
)

func TestSet(t *testing.T) {
	s := gopie.NewSet[string]()
	s.Add("hello")
	shouldEqual(t, s.Contains("hello"), true)

	s.Remove("hello")
	shouldEqual(t, s.Contains("hello"), false)
}
