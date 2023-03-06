package set

import "golang.org/x/exp/constraints"

type Set[T constraints.Ordered] struct {
	m map[T]struct{}
}

func New[T constraints.Ordered](startingSize int) *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}, startingSize),
	}
}

func (s *Set[T]) Add(e T) {
	s.m[e] = struct{}{}
}

func (s *Set[T]) Remove(e T) {
	delete(s.m, e)
}

func (s *Set[T]) Contains(e T) bool {
	_, ok := s.m[e]
	return ok
}

func (s *Set[T]) Size() int {
	return len(s.m)
}
