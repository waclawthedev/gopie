package gopie

// Set is implementation of threads unsafe set.
type Set[T comparable] interface {
	// Add adds element to set.
	Add(v T)
	// Remove removes element from set.
	Remove(v T)
	// Contains returns true if element is presented in set.
	Contains(v T) bool
	// Iterate iterates ofer Set executing passed function on every element. Once function returns false - iteration stops.
	Iterate(f func(v T) bool)
}

type threadUnsafeSet[T comparable] struct {
	m map[T]struct{}
}

// NewSet creates set for elements of provided type.
func NewSet[T comparable]() Set[T] {
	return threadUnsafeSet[T]{m: make(map[T]struct{})}
}

// Add adds element to set.
func (s threadUnsafeSet[T]) Add(v T) {
	s.m[v] = struct{}{}
}

// Remove removes element from set.
func (s threadUnsafeSet[T]) Remove(v T) {
	delete(s.m, v)
}

// Contains returns true if element is presented in set.
func (s threadUnsafeSet[T]) Contains(v T) bool {
	_, ok := s.m[v]

	return ok
}

// Iterate iterates ofer Set executing passed function on every element. Once function returns false - iteration stops.
func (s threadUnsafeSet[T]) Iterate(f func(v T) bool) {
	for k := range s.m {
		if !f(k) {
			break
		}
	}
}
