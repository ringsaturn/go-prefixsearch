// Package prefixsearch implements simple tree-based prefix search that
// i'm using for different web autocomplete services
package prefixsearch

import (
	"strings"
	"unicode"
)

// SearchTree is struct to handle search tree
type SearchTree[T comparable] struct {
	root *node[T]
}

type node[T comparable] struct {
	value    T
	childnum uint
	childs   map[rune]*node[T]
}

// New creates new search tree
func New[T comparable]() *SearchTree[T] {
	return &SearchTree[T]{
		root: &node[T]{childs: map[rune]*node[T]{}},
	}
}

// Add one leaf to tree
func (tree *SearchTree[T]) Add(key string, value T) {
	current := tree.root

	needUpdate := (*new(T) == tree.Search(key))

	for _, sym := range strings.ToLower(key) {
		if needUpdate {
			current.childnum++
		}
		next, ok := current.childs[sym]
		if !ok {
			newone := &node[T]{childs: map[rune]*node[T]{}}
			current.childs[sym] = newone
			next = newone
		}
		current = next
	}

	if needUpdate {
		current.childnum++
	}
	current.value = value
}

// AutoComplete returns autocomplete suggestions for given prefix
func (tree *SearchTree[T]) AutoComplete(prefix string) []T {
	// walk thru prefix symbols
	current := tree.root
	for _, sym := range prefix {
		var ok bool
		current, ok = current.childs[unicode.ToLower(sym)]
		if !ok {
			return make([]T, 0)
		}
	}

	// we have found, now very stupid tree walk :)
	result := make([]T, 0, current.childnum)
	current.recurse(func(v T) {
		if *new(T) != v {
			result = append(result, v)
		}
	})
	return result
}

// Search searches for value of key
func (tree *SearchTree[T]) Search(key string) T {
	current := tree.root
	for _, sym := range key {
		var ok bool
		current, ok = current.childs[unicode.ToLower(sym)]
		if !ok {
			return *new(T)
		}
	}
	return current.value
}

func (n *node[T]) recurse(callback func(T)) {
	callback(n.value)
	for _, v := range n.childs {
		v.recurse(callback)
	}
}
