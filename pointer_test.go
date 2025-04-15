package gopie_test

import (
	"gopie"
	"testing"
)

func TestPtrTo(t *testing.T) {
	result := gopie.PtrTo("123")
	shouldEqual(t, *result, "123")
}

func TestPtrValueWithDefault(t *testing.T) {
	p := gopie.PtrTo(123)
	shouldEqual(t, gopie.PtrValueOrDefault(p, 0), 123)
	shouldEqual(t, gopie.PtrValueOrDefault(nil, 111), 111)
}

func TestPtrValueOrZero(t *testing.T) {
	p := gopie.PtrTo(123)

	shouldEqual(t, gopie.PtrValueOrZero(p), 123)

	var stringPtr *string

	shouldEqual(t, gopie.PtrValueOrZero(stringPtr), "")

	var intPtr *int

	shouldEqual(t, gopie.PtrValueOrZero(intPtr), 0)
}
