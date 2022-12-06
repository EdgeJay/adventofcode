package data

import (
	"fmt"

	"github.com/thoas/go-funk"
)

type Stack[V any] struct {
	Items []V
}

func NewStack[V any]() *Stack[V] {
	return &Stack[V]{Items: []V{}}
}

func (s *Stack[V]) Add(item V) *Stack[V] {
	s.Items = append(s.Items, item)
	return s
}

func (s *Stack[V]) AddItems(items []V) *Stack[V] {
	s.Items = append(s.Items, items...)
	return s
}

func (s *Stack[V]) RemoveItems(count int, reverse bool) []V {
	if count > len(s.Items) {
		return nil
	}

	removed := s.Items[len(s.Items)-count:]
	updated := s.Items[0 : len(s.Items)-count]
	s.Items = updated

	if reverse {
		return funk.Reverse(removed).([]V)
	}

	return removed
}

func (s *Stack[V]) ReverseItems() *Stack[V] {
	s.Items = funk.Reverse(s.Items).([]V)
	return s
}

func (s *Stack[V]) LastItem() V {
	return s.Items[len(s.Items)-1]
}

func (s *Stack[V]) Print() string {
	return fmt.Sprintf("%v", s.Items)
}
