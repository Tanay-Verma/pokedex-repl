package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mu       *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cache:    make(map[string]cacheEntry),
		mu:       &sync.RWMutex{},
		interval: interval,
	}
	go newCache.readLoop()
	return newCache
}

func (c *Cache) Add(key string, value []byte) {
	if value == nil {
		fmt.Println("Nil Value not allowed")
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	cacheEntry, ok := c.cache[key]
	return cacheEntry.val, ok
}

func (c *Cache) readLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for key, ce := range c.cache {
				if ce.createdAt.Add(c.interval).Before(time.Now()) {
					delete(c.cache, key)
				}
			}
			c.mu.Unlock()
		}
	}
}
