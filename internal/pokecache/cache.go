package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}


type Cache struct {
	entries map[string]cacheEntry
	interval time.Duration
	mu *sync.Mutex
}

func NewCache (interval time.Duration) *Cache {
	cache := Cache{
		interval: interval,
		entries: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}
	go cache.reapLoop()
	return &cache
}

func (c *Cache) Add (key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.entries[key] = entry
}

func (c *Cache) Get (key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.entries[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop () {
	ticker := time.NewTicker(c.interval)
	for _ = range ticker.C {
		c.mu.Lock()
		for key := range c.entries {
			if time.Since(c.entries[key].createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
