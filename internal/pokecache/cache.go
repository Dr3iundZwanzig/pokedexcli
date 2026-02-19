package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entryMap map[string]cacheEntry
	mu       *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entryMap: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entryMap[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}

	fmt.Println("Added to cache")
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.entryMap[key]
	return value.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entryMap {
		if entry.createdAt.Before(now.Add(-last)) {
			delete(c.entryMap, key)
		}
	}
}
