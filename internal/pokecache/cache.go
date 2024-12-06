package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Calls map[string]cacheEntry
	mu    sync.RWMutex
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	c.Calls[key] = cacheEntry{time.Now(), value}
	c.mu.Unlock()
}

func NewCache(retainSeconds int) *Cache {
	c := Cache{}
	go c.sweepLoop(retainSeconds)
	return &c
}

func (c *Cache) Get(key string) (value []byte, success bool) {
	c.mu.RLock()
	entry, success := c.Calls[key]
	if success {
		return entry.val, success
	}
	return nil, success
}

func (c *Cache) sweepLoop(retainSeconds int) {
	ticker := time.NewTicker(time.Duration(retainSeconds) * time.Second)
	for {
		c.mu.Lock()
		for key, value := range c.Calls {
			if time.Until(value.createdAt)*-1 > time.Duration(retainSeconds)*time.Second {
				delete(c.Calls, key)
			}
		}
		c.mu.Unlock()
		<-ticker.C
	}
}
