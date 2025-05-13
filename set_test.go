package gopie_test

import (
	"testing"

	"github.com/waclawthedev/gopie"
)

func TestSet(t *testing.T) {
	s := gopie.NewSet[string]()
	s.Add("hello")
	shouldEqual(t, s.Contains("hello"), true)

	s.Remove("hello")
	shouldEqual(t, s.Contains("hello"), false)

	s.Add("one")
	s.Add("two")
	s.Add("three")
	s.Add("four")

	var twoWords []string

	s.Iterate(func(v string) bool {
		if len(twoWords) >= 2 {
			return false
		}

		twoWords = append(twoWords, v)

		return true
	})
	shouldEqual(t, len(twoWords), 2)
}
