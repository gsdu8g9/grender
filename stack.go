package main

import (
	"path/filepath"
)

// Stack stores a set of keyed map[string]interface{} metadata.
// The key is specified during Add, and is assumed to be a valid
// file path. Get supplies a path, and returns the union of all
// Added metadata along every step in the given path.
//
// For example, Get("/foo/bar/baz") returns merged metadata for "/foo",
// "/foo/bar", and "/foo/bar/baz". In this way, the Stack enables the
// 'stackable' Grender context behavior.
type Stack struct {
	m map[string]map[string]interface{} // path: partial-metadata
}

func NewStack() *Stack {
	return &Stack{
		m: map[string]map[string]interface{}{},
	}
}

func (s *Stack) Add(path string, m map[string]interface{}) {
	key := filepath.Join(splitPath(path)...)
	s.m[key] = m
}

func (s *Stack) Get(path string) map[string]interface{} {
	m := map[string]interface{}{}
	list := splitPath(path)
	if len(list) <= 0 {
		return m
	}

	for i, _ := range list {
		key := filepath.Join(list[:i+1]...)
		if m0, ok := s.m[key]; ok {
			m = mergeInto(m, m0)
		}
	}
	return m
}

// mergeInto merges the src map into the dst map, returning the union.
// Key collisions are handled by preferring src.
func mergeInto(dst, src map[string]interface{}) map[string]interface{} {
	for k, v := range src {
		dst[k] = v
	}
	return dst
}