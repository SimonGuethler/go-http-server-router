package router

import "strings"

type TrieNode struct {
	children map[string]*TrieNode
	route    *Route
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

func (t *Trie) Insert(route *Route) {
	pathParts := splitPath(route.path)

	if len(pathParts) > 0 {
		insertNode(t.root, pathParts, route)
	} else {
		t.root.route = route
	}
}

func insertNode(node *TrieNode, pathParts []string, route *Route) {
	if len(pathParts) == 0 {
		node.route = route
		return
	}

	first := pathParts[0]
	pathParts = pathParts[1:]

	value, ok := node.children[first]

	if ok {
		insertNode(value, pathParts, route)
	} else {
		newNode := NewTrieNode()
		node.children[first] = newNode
		insertNode(newNode, pathParts, route)
	}
}

func (t *Trie) Search(path string) *Route {
	pathParts := splitPath(path)
	return searchNode(t.root, pathParts)
}

func searchNode(node *TrieNode, pathParts []string) *Route {
	if len(pathParts) == 0 {
		return node.route
	}

	first := pathParts[0]
	pathParts = pathParts[1:]

	value, ok := node.children[first]

	if ok {
		return searchNode(value, pathParts)
	}

	return nil
}

func splitPath(path string) []string {
	path = strings.Trim(path, "/")
	pathParts := strings.Split(path, "/")
	return pathParts
}
