package router

import (
	"fmt"
	"testing"
)

func BenchmarkTrie_Insert(b *testing.B) {
	trie := NewTrie()

	for i := 0; i < b.N; i++ {
		// Generate unique routes for each iteration
		route := &Route{
			path:   fmt.Sprintf("/route/%d", i),
			method: GET,
		}
		if err := trie.Insert(route); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTrie_Search(b *testing.B) {
	trie := NewTrie()
	// Insert a sample route
	for i := 0; i < 1000; i++ {
		route := &Route{
			path:   fmt.Sprintf("/route/%d", i),
			method: GET,
		}
		if err := trie.Insert(route); err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Random route search
		_, _ = trie.Search(fmt.Sprintf("/route/%d", i%1000), GET)
	}
}
