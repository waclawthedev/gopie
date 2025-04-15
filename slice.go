package gopie

// Map returns array with every element of s passed through the mapFunc.
func Map[K any](s []K, mapFunc func(K) K) []K {
	var result = make([]K, len(s))

	for i, elem := range s {
		result[i] = mapFunc(elem)
	}

	return result
}

// Filter returns array with elements of s if filterFunc returned true for according element.
func Filter[K any](s []K, filterFunc func(K) bool) []K {
	var result = make([]K, 0)

	for _, elem := range s {
		if filterFunc(elem) {
			result = append(result, elem)
		}
	}

	return result
}

// Reduce returns one K var that is result of cumulative execution of reduceFunc on every element of s.
func Reduce[K any](s []K, reduceFunc func(K, K) K) (result K) {
	if len(s) == 0 {
		return
	}

	result = s[0]

	if len(s) == 1 {
		return
	}

	for i := 1; i < len(s); i++ {
		result = reduceFunc(result, s[i])
	}

	return
}

// Flatten returns slice based on matrix values.
func Flatten[K any](s [][]K) (result []K) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			result = append(result, s[i][j])
		}
	}

	return
}

// SplitByN splits a slice s into sub-slices of maximum size n.
// Returns nil if n is less than or equal to 0.
func SplitByN[T any](s []T, n int) [][]T {
	if n <= 0 || len(s) == 0 {
		return nil // Invalid chunk size or empty input slice
	}

	// Calculate the number of chunks needed, rounded up
	count := (len(s) + n - 1) / n
	result := make([][]T, count)

	for i := 0; i < count; i++ {
		start := i * n
		end := start + n

		// Ensure the last chunk doesn't exceed slice bounds
		if end > len(s) {
			end = len(s)
		}

		result[i] = s[start:end]
	}

	return result
}
