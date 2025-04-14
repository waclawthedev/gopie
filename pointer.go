package gopie

// PtrTo return pointer with value passed.
func PtrTo[T any](v T) *T {
	return &v
}

// PtrValueOrDefault safely returns value by pointer or returns default value.
func PtrValueOrDefault[T any](v *T, d T) T {
	if v != nil {
		return *v
	}

	return d
}

// PtrValueOrZero returns value by pointer or zero value according to type.
func PtrValueOrZero[T any](v *T) T {
	if v != nil {
		return *v
	}

	var zero T

	return zero
}
