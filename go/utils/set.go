package utils

type Set[T comparable] struct {
	set map[T]bool
}

func NewSet[T comparable]() (newSet Set[T]) {
	newSet.set = make(map[T]bool)
	return
}

func (s *Set[T]) Add(elements ...T) {
	for _, element := range elements {
		s.set[element] = true
	}
}

func (s *Set[T]) Remove(elements ...T) (ok bool) {
	for _, element := range elements {
		if !s.Exists(element) {
			ok = false
			return
		}
	}
	for _, element := range elements {
		delete(s.set, element)
	}
	ok = true
	return
}

func (s *Set[T]) Exists(element T) bool {
	return s.set[element]
}
