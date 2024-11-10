package router

import (
	"reflect"
	"testing"
)

func TestSanitizePath(t *testing.T) {
	tests := []struct {
		path         string
		allowDynamic bool
		expected     string
		expectError  bool
	}{
		{"/valid/path", false, "/valid/path", false},
		{"//extra/slashes//", false, "/extra/slashes", false}, // extra slashes
		{"/:dynamic/path", true, "/:dynamic/path", false},     // dynamic allowed
		{"/:dynamic/path", false, "", true},                   // dynamic not allowed
		{"", false, "/", false},                               // empty path
		{"/", false, "/", false},                              // root path
		{"invalid@path", false, "", true},                     // invalid characters
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			result, err := SanitizePath(tt.path, tt.allowDynamic)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error for path %s, got nil", tt.path)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error for path %s: %v", tt.path, err)
				}
				if result != tt.expected {
					t.Errorf("expected %s, got %s", tt.expected, result)
				}
			}
		})
	}
}

func TestSanitizePathElement(t *testing.T) {
	tests := []struct {
		element     string
		pattern     string
		expected    string
		expectError bool
	}{
		{"valid", pathElementPattern, "valid", false},
		{":dynamic", dynamicPathAllowedChars, ":dynamic", false},
		{"invalid@", pathElementPattern, "", true}, // invalid character
		{"another-valid", pathElementPattern, "another-valid", false},
	}

	for _, tt := range tests {
		t.Run(tt.element, func(t *testing.T) {
			result, err := SanitizePathElement(tt.element, tt.pattern)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error for element %s, got nil", tt.element)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error for element %s: %v", tt.element, err)
				}
				if result != tt.expected {
					t.Errorf("expected %s, got %s", tt.expected, result)
				}
			}
		})
	}
}

func TestIsPathParam(t *testing.T) {
	tests := []struct {
		segment  string
		expected bool
	}{
		{":param", true},           // valid path param
		{":valid-param", true},     // valid param with hyphen
		{"path", false},            // static path
		{": invalid param", false}, // invalid due to space
	}

	for _, test := range tests {
		t.Run(test.segment, func(t *testing.T) {
			result := IsPathParam(test.segment)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestExtractPathVariables(t *testing.T) {
	tests := []struct {
		routePath    string
		requestPath  string
		expectedVars map[string]string
	}{
		{"/users/:id", "/users/123", map[string]string{"id": "123"}},
		{"/users/:userId/books/:bookId", "/users/45/books/789", map[string]string{"userId": "45", "bookId": "789"}},
		{"/:category/:id", "/electronics/567", map[string]string{"category": "electronics", "id": "567"}},
		{"/products/:productId/reviews", "/products/42/reviews", map[string]string{"productId": "42"}},
		{"/static/path", "/static/path", map[string]string{}}, // no params
		{"/books/:id/authors/:authorId", "/books/5/authors/12", map[string]string{"id": "5", "authorId": "12"}},
		{"/multiple/:var1/segments/:var2", "/multiple/one/segments/two", map[string]string{"var1": "one", "var2": "two"}},
	}

	for _, tt := range tests {
		t.Run(tt.routePath, func(t *testing.T) {
			result := ExtractPathVariables(tt.routePath, tt.requestPath)
			if !reflect.DeepEqual(result, tt.expectedVars) {
				t.Errorf("for route %s and request %s, expected %v, got %v", tt.routePath, tt.requestPath, tt.expectedVars, result)
			}
		})
	}
}

func TestSplitPath(t *testing.T) {
	tests := []struct {
		path     string
		expected []string
	}{
		{"", []string{""}},  // edge case empty path
		{"/", []string{""}}, // root path
		{"/path/to/resource", []string{"path", "to", "resource"}},
		{"path/to/resource/", []string{"path", "to", "resource"}}, // trailing slash
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			result := SplitPath(tt.path)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsPathParam3(t *testing.T) {
	tests := []struct {
		pathSegment string
		expected    bool
	}{
		{":dynamic", true},  // valid dynamic param
		{"static", false},   // static path
		{":123param", true}, // valid param with numbers
		{":-invalid", true}, // valid param with hyphen
		{":", false},        // invalid due to missing name
		{":a", true},        // valid single char param
	}

	for _, tt := range tests {
		t.Run(tt.pathSegment, func(t *testing.T) {
			result := IsPathParam(tt.pathSegment)
			if result != tt.expected {
				t.Errorf("for segment %s, expected %v, got %v", tt.pathSegment, tt.expected, result)
			}
		})
	}
}
