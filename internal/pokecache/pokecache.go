package pokeapi

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
		mu:      sync.Mutex{},
	}
	cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.entries[key]
	return v.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				c.mu.Lock()
				now := time.Now()
				for k, v := range c.entries {
					if now.Sub(v.createdAt) > interval {
						delete(c.entries, k)
					}
				}
				c.mu.Unlock()
			}
		}
	}()
	// c.mu.Lock()
	// defer c.mu.Unlock()
	// for k, v := range c.cache {
	// }

}
