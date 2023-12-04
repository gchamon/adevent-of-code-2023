package utils

type Set[T comparable] struct {
	Map map[T]bool
}

func NewSet[T comparable](elements ...T) (newSet Set[T]) {
	newSet.Map = make(map[T]bool)
	if len(elements) > 0 {
		newSet.Add(elements...)
	}
	return
}

func (s *Set[T]) Add(elements ...T) {
	for _, element := range elements {
		s.Map[element] = true
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
		delete(s.Map, element)
	}
	ok = true
	return
}

func (s *Set[T]) Exists(element T) bool {
	return s.Map[element]
}

func (s *Set[T]) Len() int {
	return len(s.Map)
}

func (s Set[T]) Union(si ...Set[T]) (u Set[T]) {
	u = NewSet[T]()
	for _, set := range append(si, s) {
		for k := range set.Map {
			u.Add(k)
		}
	}
	return
}

func intersect[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	i := NewSet[T]()
	var smaller, larger Set[T]
	if s1.Len() < s2.Len() {
		smaller, larger = s1, s2
	} else {
		smaller, larger = s2, s1
	}
	for k := range smaller.Map {
		if larger.Exists(k) {
			i.Add(k)
		}
	}
	return i
}

func (s Set[T]) Intersection(si ...Set[T]) Set[T] {
	i := intersect(s, si[0])
	for _, current := range si[1:] {
		i = intersect(i, current)
	}
	return i
}
