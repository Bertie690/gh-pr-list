package utils

// A Set is a basic implementation of a key-only hashmap.
type Set[T comparable] map[T]struct{}

// NewSet creates a new Set containing the provided items.
func NewSet[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	for _, item := range items {
		s.Add(item)
	}
	return s
}

// Contains reports whether value v is inside set s.
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// Add inserts value v into set s.
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}
