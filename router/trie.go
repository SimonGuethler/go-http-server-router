package router

import (
	"errors"
	"strings"
)

type TrieNode struct {
	children map[string]*TrieNode
	route    map[HTTPMethod]*Route
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[string]*TrieNode)}
}

func (t *Trie) Insert(route *Route) error {
	path, err := SanitizePath(route.path, true)
	if err != nil {
		return err
	}
	pathParts := SplitPath(path)

	if len(pathParts) > 0 {
		err := insertNode(t.root, pathParts, route)
		if err != nil {
			return err
		}
	} else {
		if t.root.route != nil {
			if _, ok := t.root.route[route.method]; ok {
				return errors.New("route already exists")
			} else {
				t.root.route[route.method] = route
			}
		} else {
			t.root.route = make(map[HTTPMethod]*Route)
			t.root.route[route.method] = route
		}
	}

	return nil
}

func insertNode(node *TrieNode, pathParts []string, route *Route) error {
	if len(pathParts) == 0 {
		if node.route != nil {
			if _, ok := node.route[route.method]; ok {
				return errors.New("route already exists")
			} else {
				node.route[route.method] = route
			}
		} else {
			node.route = make(map[HTTPMethod]*Route)
			node.route[route.method] = route
		}
		return nil
	}

	first := pathParts[0]
	pathParts = pathParts[1:]

	isParam := IsPathParam(first)
	if isParam {
		first = "*"
	} else {
		first = strings.ToLower(first)
	}

	value, ok := node.children[first]

	if ok {
		err := insertNode(value, pathParts, route)
		if err != nil {
			return err
		}
	} else {
		newNode := NewTrieNode()
		node.children[first] = newNode
		err := insertNode(newNode, pathParts, route)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Trie) Search(path string, method HTTPMethod) (*Route, error) {
	path, err := SanitizePath(path, false)
	if err != nil {
		return nil, err
	}
	pathParts := SplitPath(path)
	return searchNode(t.root, pathParts, method), nil
}

func searchNode(node *TrieNode, pathParts []string, method HTTPMethod) *Route {
	if len(pathParts) == 0 {
		if node.route != nil {
			if route, ok := node.route[method]; ok {
				return route
			}
		}

		return nil
	}

	first := pathParts[0]
	pathParts = pathParts[1:]

	value, ok := node.children[first]

	if ok {
		return searchNode(value, pathParts, method)
	} else {
		value, ok := node.children["*"]
		if ok {
			return searchNode(value, pathParts, method)
		}
	}

	return nil
}
