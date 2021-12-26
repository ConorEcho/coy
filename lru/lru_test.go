package lru

import "testing"

func TestAdd(t *testing.T) {
	cache := New(2)

	cache.Add("foo", 1)

	if used := cache.Used(); used != 1 {
		t.Fatalf("cache used=%d, want=1", used)
	}

	// add repeatedly will not add the curSize
	cache.Add("foo", 1)
	if used := cache.Used(); used != 1 {
		t.Fatalf("cache used=%d, want=1", used)
	}

	cache.Add("bar", 1)
	if used := cache.Used(); used != 2 {
		t.Fatalf("cache used=%d, want=2", used)
	}

	// Cache will cache 2 elem at most.
	cache.Add("key", 1)
	if used := cache.Used(); used != 2 {
		t.Fatalf("cache used=%d, want=2", used)
	}
}

func TestGet(t *testing.T) {
	cache := New(1)
	cache.Add("foo", "1")

	if val, ok := cache.Get("foo"); !ok || val.(string) != "1" {
		t.Fatalf("cache hit foo=1 failed")
	}

	if _, ok := cache.Get("bar"); ok {
		t.Fatalf("cache miss bar failed")
	}

	cache.Add("bar", 2)

	// because cache size is 1, so foo should be evicted.
	if _, ok := cache.Get("foo"); ok {
		t.Fatalf("cache miss foo failed")
	}

	if val, ok := cache.Get("bar"); !ok || val.(int) != 2 {
		t.Fatalf("cache hit bar=2 failed")
	}
}
