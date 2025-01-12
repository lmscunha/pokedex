package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           sync.Mutex
	stop         chan struct{}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		stop:         make(chan struct{}),
	}

	if interval != 0 {
		go cache.reapLoop(interval)
	}

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheEntries[key]
	if ok == false {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-c.stop:
			return
		case <-ticker.C:
			c.mu.Lock()
			now := time.Now()
			for k, entry := range c.cacheEntries {
				if now.Sub(entry.createdAt) >= interval {
					delete(c.cacheEntries, k)
				}
			}
			c.mu.Unlock()
		}
	}
}

func (c *Cache) Stop() {
	close(c.stop)
}
