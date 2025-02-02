package pokeapi

import (
	"testing"
	"time"
)

// Personnal test
func TestCache(t *testing.T) {
	cache := NewCache(5 * time.Second)

	cases := struct {
		input *Cache
		// expectd []string
	}{
		input: cache,
	}

	cases.input.Add("foo", []byte("foobar"))
	b, ok := cache.Get("foo")
	if !ok {
		t.Errorf("foo doesn't exist!\n")
	}
	t.Logf("before: %s | %v\n", b, ok)
	time.Sleep(6 * time.Second)
	b, ok = cases.input.Get("foo")
	if ok {
		t.Errorf("foo should not exist!\n")
	}
	t.Logf("after: %s | %v\n", b, ok)
}

// Given by Boots
func TestCacheBoots(t *testing.T) {
	cache := NewCache(5 * time.Second)

	// Test adding and immediate retrieval
	cache.Add("foo", []byte("foobar"))
	val, ok := cache.Get("foo")
	if !ok {
		t.Error("expected to find key immediately after adding")
	}
	if string(val) != "foobar" {
		t.Errorf("expected 'foobar', got '%s'", string(val))
	}

	// Test reaping
	time.Sleep(6 * time.Second)
	_, ok = cache.Get("foo")
	if ok {
		t.Error("expected key to be reaped after interval")
	}
}
