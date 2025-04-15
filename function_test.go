package gopie_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/waclawthedev/gopie"
)

func TestMust(t *testing.T) {
	shouldEqual(t, string(gopie.Must(json.Marshal(struct {
		Field string `json:"field"`
	}{
		Field: "value",
	}))), `{"field":"value"}`)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the Must function did not panic")
		}
	}()

	gopie.Must(json.Marshal(func() {})) //nolint:errchkjson,staticcheck
}

func ExampleMust() {
	fmt.Println(string(gopie.Must(json.Marshal("hello world"))))
	// Output: "hello world"
}
