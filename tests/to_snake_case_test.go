package tests

import (
	"testing"

	"github.com/ebina4yaka/gqlgen-api-example/lib"
)

func TestToSnakeCase(t *testing.T) {
	var actual string
	var expected string

	// from camelCase
	actual = lib.ToSnakeCase("createdAt")
	expected = "created_at"
	if actual != expected {
		t.Errorf("got %s\nwant %s", actual, expected)
	}

	// from PascalCase
	actual = lib.ToSnakeCase("UserProfile")
	expected = "user_profile"
	if actual != expected {
		t.Errorf("got %s\nwant %s", actual, expected)
	}

	// from snake_case (don't convert)
	actual = lib.ToSnakeCase("post_user")
	expected = "post_user"
	if actual != expected {
		t.Errorf("got %s\nwant %s", actual, expected)
	}
}

func BenchmarkToSnakeCase(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lib.ToSnakeCase("BenchmarkToSnakeCase")
	}
}
