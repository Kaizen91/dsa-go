package lru

import (
	"testing"
)

func TestLRU(t *testing.T) {
	cache := New[string, int](2)
	cache.Put("foo", 69)
	cache.Put("bar", 420)
	if _, ok := cache.Get("foo"); !ok {
		t.Fatalf("Expected key 'foo' to be present")
	}
	cache.Put("pie", 314)
	if _, ok := cache.Get("bar"); ok {
		t.Errorf("Expected key 'bar' to be evicted, but it was found")
	}
	if val, ok := cache.Get("foo"); !ok || val != 69 {
		t.Errorf("Expected foo = 69, got foo = %v (found: %v)", val, ok)
	}
	if val, ok := cache.Get("pie"); !ok || val != 314 {
		t.Errorf("Expected pie = 314, got pie = %v (found: %v)", val, ok)
	}
}
