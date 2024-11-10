package router

import (
	"reflect"
	"testing"
)

// Unit Tests

func TestSanitizePath(t *testing.T) {
	tests := []struct {
		path      string
		expected  string
		expectErr bool
	}{
		{"/path/to/something", "/path/to/something", false},
		{"path/to/something", "/path/to/something", false},
		{"/invalid_path!@#", "", true},          // invalid path element
		{"//path//to///something///", "", true}, // extra slashes
		{"/", "/", false},
		{"", "/", false},
		{"asdf", "/asdf", false},
		{"/asdf/", "/asdf", false},
		{"asdf/", "/asdf", false},
		{"///asdf///", "/asdf", false},
		{"///asdf///qwer///", "", true},
	}

	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			result, err := SanitizePath(test.path, false)
			if test.expectErr {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != test.expected {
					t.Errorf("expected %s, got %s", test.expected, result)
				}
			}
		})
	}
}

func TestSanitizePathElement(t *testing.T) {
	tests := []struct {
		element   string
		expected  string
		expectErr bool
	}{
		{"valid_element", "valid_element", false},
		{"INVALID!@", "", true}, // invalid element
		{"another_valid-elem", "another_valid-elem", false},
	}

	for _, test := range tests {
		t.Run(test.element, func(t *testing.T) {
			result, err := SanitizePathElement(test.element, pathElementPattern)
			if test.expectErr {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != test.expected {
					t.Errorf("expected %s, got %s", test.expected, result)
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
		{":param", true},
		{":valid-param", true},
		{"path", false},
		{":invalid param", false},
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

// Benchmarks

func BenchmarkSanitizePath(b *testing.B) {
	tests := []struct {
		path string
	}{
		{"/path/to/something"},
		{"path/to/something"},
		{"/path/to//something/with///extra/slashes"},
	}

	for _, test := range tests {
		b.Run(test.path, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = SanitizePath(test.path, false)
			}
		})
	}
}

func BenchmarkSanitizePathElement(b *testing.B) {
	tests := []struct {
		element string
	}{
		{"valid_element"},
		{"INVALID!@"}, // Invalid
		{"another_valid-elem"},
	}

	for _, test := range tests {
		b.Run(test.element, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = SanitizePathElement(test.element, pathElementPattern)
			}
		})
	}
}

func BenchmarkSplitPath(b *testing.B) {
	tests := []struct {
		path string
	}{
		{"/path/to/something"},
		{"/another/example/path"},
		{"path"},
		{"/longer/path/to/another/example"},
	}

	for _, test := range tests {
		b.Run(test.path, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = SplitPath(test.path)
			}
		})
	}
}

func TestExtractPathVariables(t *testing.T) {
	tests := []struct {
		routePath   string
		requestPath string
		expected    map[string]string
	}{
		{
			routePath:   "/users/:id",
			requestPath: "/users/123",
			expected:    map[string]string{"id": "123"},
		},
		{
			routePath:   "/users/:userId/books/:bookId",
			requestPath: "/users/45/books/789",
			expected:    map[string]string{"userId": "45", "bookId": "789"},
		},
		{
			routePath:   "/:category/:id",
			requestPath: "/electronics/567",
			expected:    map[string]string{"category": "electronics", "id": "567"},
		},
		{
			routePath:   "/products/:productId/reviews",
			requestPath: "/products/42/reviews",
			expected:    map[string]string{"productId": "42"},
		},
		{
			routePath:   "/fixed/path",
			requestPath: "/fixed/path",
			expected:    map[string]string{},
		},
		{
			routePath:   "/books/:id/authors/:authorId",
			requestPath: "/books/5/authors/12",
			expected:    map[string]string{"id": "5", "authorId": "12"},
		},
		{
			routePath:   "/multiple/:var1/segments/:var2",
			requestPath: "/multiple/one/segments/two",
			expected:    map[string]string{"var1": "one", "var2": "two"},
		},
	}

	for _, test := range tests {
		result := ExtractPathVariables(test.routePath, test.requestPath)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For routePath '%s' and requestPath '%s', expected %v, but got %v",
				test.routePath, test.requestPath, test.expected, result)
		}
	}
}
