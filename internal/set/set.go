package set

import (
	"github.com/WadeGulbrandsen/aoc2023/internal/utils"
)

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) Delete(item T) {
	delete(s, item)
}

func (s Set[T]) Contains(item T) bool {
	_, ok := (s)[item]
	return ok
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) Intersect(o Set[T]) Set[T] {
	n := NewSet[T]()
	for k := range s {
		if o.Contains(k) {
			n.Add(k)
		}
	}
	return n
}

func (s Set[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}
	item := s.Slice()[0]
	s.Delete(item)
	return item, true
}

func (s Set[T]) Slice() []T {
	return utils.GetMapKeys(s)
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}
