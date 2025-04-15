// Package gopie is main package of that module
package gopie

// Must function allows to skip err check if it is ok for you to call panic in case of err!=nil.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}
