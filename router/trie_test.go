package router

import (
	"testing"
)

func TestInsertAndSearchLongPaths(t *testing.T) {
	// Create a new Trie instance
	trie := NewTrie()

	// Define longer static routes for testing
	route1 := &Route{path: "/users/profile/details", method: "GET"}
	route2 := &Route{path: "/products/category/electronics/item/123", method: "POST"}
	route3 := &Route{path: "/admin/dashboard/settings", method: "PUT"}

	// Insert longer routes into the trie
	err := trie.Insert(route1)
	if err != nil {
		t.Errorf("Insert failed for route1: %v", err)
	}

	err = trie.Insert(route2)
	if err != nil {
		t.Errorf("Insert failed for route2: %v", err)
	}

	err = trie.Insert(route3)
	if err != nil {
		t.Errorf("Insert failed for route3: %v", err)
	}

	// Test searching for the routes
	searchResult, _ := trie.Search("/users/profile/details", "GET")
	if searchResult == nil {
		t.Error("Search failed for route1")
	} else if searchResult != route1 {
		t.Errorf("Search result for '/users/profile/details' should be route1, got %v", searchResult)
	}

	searchResult, _ = trie.Search("/products/category/electronics/item/123", "POST")
	if searchResult == nil {
		t.Error("Search failed for route2")
	} else if searchResult != route2 {
		t.Errorf("Search result for '/products/category/electronics/item/123' should be route2, got %v", searchResult)
	}

	searchResult, _ = trie.Search("/admin/dashboard/settings", "PUT")
	if searchResult == nil {
		t.Error("Search failed for route3")
	} else if searchResult != route3 {
		t.Errorf("Search result for '/admin/dashboard/settings' should be route3, got %v", searchResult)
	}
}

func TestInsertAndSearchLongPathsWithDifferentMethods(t *testing.T) {
	// Create a new Trie instance
	trie := NewTrie()

	// Define some routes with longer paths and different methods
	route1 := &Route{path: "/users/profile/details", method: "GET"}
	route2 := &Route{path: "/users/profile/details", method: "POST"}
	route3 := &Route{path: "/admin/dashboard/settings", method: "PUT"}

	// Insert routes into the trie
	err := trie.Insert(route1)
	if err != nil {
		t.Errorf("Insert failed for route1: %v", err)
	}

	err = trie.Insert(route2)
	if err != nil {
		t.Errorf("Insert failed for route2: %v", err)
	}

	err = trie.Insert(route3)
	if err != nil {
		t.Errorf("Insert failed for route3: %v", err)
	}

	// Test searching for the routes with different methods
	searchResult, _ := trie.Search("/users/profile/details", "GET")
	if searchResult == nil {
		t.Error("Search failed for route1 (GET)")
	} else if searchResult != route1 {
		t.Errorf("Search result for '/users/profile/details' GET should be route1, got %v", searchResult)
	}

	searchResult, _ = trie.Search("/users/profile/details", "POST")
	if searchResult == nil {
		t.Error("Search failed for route2 (POST)")
	} else if searchResult != route2 {
		t.Errorf("Search result for '/users/profile/details' POST should be route2, got %v", searchResult)
	}

	searchResult, _ = trie.Search("/admin/dashboard/settings", "PUT")
	if searchResult == nil {
		t.Error("Search failed for route3 (PUT)")
	} else if searchResult != route3 {
		t.Errorf("Search result for '/admin/dashboard/settings' PUT should be route3, got %v", searchResult)
	}
}

func TestInsertDuplicateLongPaths(t *testing.T) {
	// Create a new Trie instance
	trie := NewTrie()

	// Define a long static route
	route1 := &Route{path: "/users/profile/details", method: "GET"}

	// Insert the first route
	err := trie.Insert(route1)
	if err != nil {
		t.Errorf("Insert failed for route1: %v", err)
	}

	// Try inserting a duplicate route (same path and method)
	err = trie.Insert(route1)
	if err == nil {
		t.Error("Insert should return an error for duplicate route")
	}

	// Insert another route with a different method for the same path
	route2 := &Route{path: "/users/profile/details", method: "POST"}
	err = trie.Insert(route2)
	if err != nil {
		t.Errorf("Insert failed for route2: %v", err)
	}

	// Test that both routes exist
	searchResult, _ := trie.Search("/users/profile/details", "GET")
	if searchResult == nil || searchResult != route1 {
		t.Errorf("Search should return route1 for GET, got %v", searchResult)
	}

	searchResult, _ = trie.Search("/users/profile/details", "POST")
	if searchResult == nil || searchResult != route2 {
		t.Errorf("Search should return route2 for POST, got %v", searchResult)
	}
}

func TestSearchLongPathNotFound(t *testing.T) {
	// Create a new Trie instance
	trie := NewTrie()

	// Define a static route with a long path
	route1 := &Route{path: "/users/profile/details", method: "GET"}
	err := trie.Insert(route1)
	if err != nil {
		t.Errorf("Insert failed for route1: %v", err)
	}

	// Test searching for a non-existent long path
	searchResult, _ := trie.Search("/nonexistent/long/path", "GET")
	if searchResult != nil {
		t.Errorf("Search should return nil for a non-existent route, got %v", searchResult)
	}
}

func TestLongPathInsertionWithTrailingSlash(t *testing.T) {
	// Create a new Trie instance
	trie := NewTrie()

	// Define a long path with a trailing slash
	route1 := &Route{path: "/users/profile/details/", method: "GET"}

	// Insert the route
	err := trie.Insert(route1)
	if err != nil {
		t.Errorf("Insert failed for route1: %v", err)
	}

	// Test searching for the route with and without the trailing slash
	searchResult, _ := trie.Search("/users/profile/details", "GET")
	if searchResult == nil {
		t.Error("Search failed for route1 (without trailing slash)")
	} else if searchResult != route1 {
		t.Errorf("Search result for '/users/profile/details' should be route1, got %v", searchResult)
	}

	searchResult, _ = trie.Search("/users/profile/details/", "GET")
	if searchResult == nil {
		t.Error("Search failed for route1 (with trailing slash)")
	} else if searchResult != route1 {
		t.Errorf("Search result for '/users/profile/details/' should be route1, got %v", searchResult)
	}
}

// Test for root (empty path) and main page ("/")

func TestInsertAndSearchRootPaths(t *testing.T) {
	// Create a new Trie instance
	trie1 := NewTrie()
	trie2 := NewTrie()

	// Define routes for the root path and empty path
	route1 := &Route{path: "", method: "GET"}  // for the empty path ""
	route2 := &Route{path: "/", method: "GET"} // for the root path "/"

	// Insert root routes
	err := trie1.Insert(route1)
	if err != nil {
		t.Errorf("Insert failed for empty path: %v", err)
	}

	err = trie1.Insert(route2)
	if err == nil {
		t.Errorf("Insert not failed for root path: %v", err)
	}

	err = trie2.Insert(route2)
	if err != nil {
		t.Errorf("Insert failed for root path: %v", err)
	}

	err = trie2.Insert(route1)
	if err == nil {
		t.Errorf("Insert not failed for empty path: %v", err)
	}

	// Test searching for both root and empty paths
	searchResult, _ := trie1.Search("", "GET")
	if searchResult == nil {
		t.Error("Search failed for empty path ''")
	} else if searchResult != route1 {
		t.Errorf("Search result for empty path '' should be route1, got %v", searchResult)
	}

	searchResult, _ = trie2.Search("/", "GET")
	if searchResult == nil {
		t.Error("Search failed for root path '/'")
	} else if searchResult != route2 {
		t.Errorf("Search result for root path '/' should be route2, got %v", searchResult)
	}
}

// Benchmark tests

func BenchmarkInsertLongPath(b *testing.B) {
	trie := NewTrie()
	route := &Route{path: "/users/profile/details", method: "GET"}

	b.ResetTimer() // To ensure the timing starts only after setup
	for i := 0; i < b.N; i++ {
		err := trie.Insert(route)
		if err != nil {
			b.Errorf("Insert failed: %v", err)
		}
	}
}

func BenchmarkSearchLongPath(b *testing.B) {
	trie := NewTrie()
	route := &Route{path: "/users/profile/details", method: "GET"}
	trie.Insert(route)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = trie.Search("/users/profile/details", "GET")
	}
}

func TestTrie_InsertAndSearch(t *testing.T) {
	// Define helper to create routes
	newRoute := func(path string, method HTTPMethod) *Route {
		return &Route{
			path:   path,
			method: method,
		}
	}

	// Test cases
	tests := []struct {
		name           string
		insertPath     string
		insertMethod   HTTPMethod
		searchPath     string
		searchMethod   HTTPMethod
		expectedResult bool
	}{
		// Test empty string path and root path "/"
		{"Empty path", "", GET, "", GET, true},
		{"Root path", "/", GET, "/", GET, true},

		// Static route tests
		{"Static route", "/static", GET, "/static", GET, true},
		{"Non-existent static route", "/static", GET, "/nonexistent", GET, false},

		// Dynamic route tests
		{"Dynamic route", "/user/:id", GET, "/user/123", GET, true},
		{"Dynamic route different path", "/user/:id", GET, "/user/456", GET, true},
		{"Dynamic route with missing part", "/user/:id", GET, "/user", GET, false},

		// Dynamic route in the middle
		{"Dynamic route in middle", "/products/:id/details", GET, "/products/123/details", GET, true},
		{"Dynamic route in middle, different ID", "/products/:id/details", GET, "/products/456/details", GET, true},
		{"Dynamic route in middle, missing part", "/products/:id/details", GET, "/products/details", GET, false},

		// Multiple dynamic routes
		{"Multiple dynamic routes", "/user/:id/orders/:orderId", GET, "/user/123/orders/456", GET, true},
		{"Multiple dynamic routes different IDs", "/user/:id/orders/:orderId", GET, "/user/789/orders/012", GET, true},
		{"Multiple dynamic routes with missing part", "/user/:id/orders/:orderId", GET, "/user/123/orders", GET, false},

		// Long routes
		{"Long static route", "/a/b/c/d/e/f", GET, "/a/b/c/d/e/f", GET, true},
		{"Long dynamic route", "/a/:b/c/:d/e/:f", GET, "/a/1/c/2/e/3", GET, true},
		{"Non-existent long route", "/a/b/c/d/e/f", GET, "/a/b/c", GET, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := NewTrie()

			route := newRoute(tt.insertPath, tt.insertMethod)
			err := trie.Insert(route)
			if err != nil {
				t.Fatalf("Failed to insert route: %v", err)
			}

			// Search for the route
			foundRoute, _ := trie.Search(tt.searchPath, tt.searchMethod)
			if (foundRoute != nil) != tt.expectedResult {
				t.Errorf("Expected route to be %v, got %v", tt.expectedResult, foundRoute != nil)
			}
		})
	}
}
