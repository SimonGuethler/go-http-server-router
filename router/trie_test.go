package router

import (
	"reflect"
	"testing"
)

func TestInsertAndSearchLongPaths(t *testing.T) {
	tests := []struct {
		name          string
		insertRoutes  []*Route
		searchPath    string
		searchMethod  HTTPMethod
		expectedRoute *Route
		expectedError bool
	}{
		{
			name: "Insert and search long static paths",
			insertRoutes: []*Route{
				{path: "/users/profile/details", method: "GET"},
				{path: "/products/category/electronics/item/123", method: "POST"},
				{path: "/admin/dashboard/settings", method: "PUT"},
			},
			searchPath:    "/users/profile/details",
			searchMethod:  "GET",
			expectedRoute: &Route{path: "/users/profile/details", method: "GET"},
		},
		{
			name: "Search for a non-existent route",
			insertRoutes: []*Route{
				{path: "/users/profile/details", method: "GET"},
			},
			searchPath:    "/nonexistent/long/path",
			searchMethod:  "GET",
			expectedRoute: nil,
		},
		{
			name: "Insert and search for root path with trailing slash",
			insertRoutes: []*Route{
				{path: "/users/profile/details/", method: "GET"},
			},
			searchPath:    "/users/profile/details",
			searchMethod:  "GET",
			expectedRoute: &Route{path: "/users/profile/details/", method: "GET"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new Trie instance
			trie := NewTrie()

			// Insert the routes
			for _, route := range tt.insertRoutes {
				if err := trie.Insert(route); err != nil {
					t.Errorf("Insert failed for route %v: %v", route, err)
				}
			}

			// Test searching for the route
			searchResult, err := trie.Search(tt.searchPath, tt.searchMethod)
			if err != nil {
				t.Errorf("Unexpected error while searching for '%s': %v", tt.searchPath, err)
			}

			// Use reflect.DeepEqual for struct comparison
			if !reflect.DeepEqual(searchResult, tt.expectedRoute) {
				t.Errorf("Search result for '%s' with method '%s' expected: %+v, got: %+v", tt.searchPath, tt.searchMethod, tt.expectedRoute, searchResult)
			}
		})
	}
}

func TestHomePathEmpty(t *testing.T) {
	trie := NewTrie()
	route := &Route{path: "", method: "GET"}

	err := trie.Insert(route)
	if err != nil {
		t.Errorf("Insert failed for route %v: %v", route, err)
	}

	searchResult, err := trie.Search("/", "GET")
	if err != nil {
		t.Errorf("Unexpected error while searching for '/': %v", err)
	}
	if !reflect.DeepEqual(searchResult, route) {
		t.Errorf("Search result for '/' with method 'GET' expected: %+v, got: %+v", route, searchResult)
	}

	searchResult, err = trie.Search("", "GET")
	if err != nil {
		t.Errorf("Unexpected error while searching for '': %v", err)
	}
	if !reflect.DeepEqual(searchResult, route) {
		t.Errorf("Search result for '' with method 'GET' expected: %+v, got: %+v", route, searchResult)
	}
}

func TestHomePathSlash(t *testing.T) {
	trie := NewTrie()
	route := &Route{path: "/", method: "GET"}

	err := trie.Insert(route)
	if err != nil {
		t.Errorf("Insert failed for route %v: %v", route, err)
	}

	searchResult, err := trie.Search("/", "GET")
	if err != nil {
		t.Errorf("Unexpected error while searching for '/': %v", err)
	}
	if !reflect.DeepEqual(searchResult, route) {
		t.Errorf("Search result for '/' with method 'GET' expected: %+v, got: %+v", route, searchResult)
	}

	searchResult, err = trie.Search("", "GET")
	if err != nil {
		t.Errorf("Unexpected error while searching for '': %v", err)
	}
	if !reflect.DeepEqual(searchResult, route) {
		t.Errorf("Search result for '/' with method 'GET' expected: %+v, got: %+v", route, searchResult)
	}
}
