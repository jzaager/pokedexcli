package pokecache

import (
	"sync"
	"time"
)

// val is the raw data being cached
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{entries: map[string]cacheEntry{}}
	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(url string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newEntry := cacheEntry{createdAt: time.Now(), val: data}
	c.entries[url] = newEntry
}

func (c *Cache) Get(url string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[url]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	// sends current time on channel each time interval passes
	// defer ticker.Stop() - no longer needed after go 1.23
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C // waits for the next tick
		c.mu.Lock()
		now := time.Now()

		for key, entry := range c.entries {
			// duration diff between now and createdAt compared to interval
			// if duration diff > interval, delete the entry
			if now.Sub(entry.createdAt) > interval {
				delete(c.entries, key) // remove old entries
			}
		}
		c.mu.Unlock()
	}
}
