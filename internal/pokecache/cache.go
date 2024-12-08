package pokecache

import (
	"sync"
	"time"
)

type cacheEntry[T any] struct {
	createdAt time.Time
	val       T
}

type Cache[T any] struct {
	Calls map[string]cacheEntry[T]
	mu    sync.RWMutex
}

func (c *Cache[T]) Add(key string, value T) {
	c.mu.Lock()
	c.Calls[key] = cacheEntry[T]{time.Now(), value}
	c.mu.Unlock()
}

func NewCache[T any](retainSeconds int) *Cache[T] {
	c := Cache[T]{}
	go c.sweepLoop(retainSeconds)
	return &c
}

func (c *Cache[T]) Get(key string) (value T, success bool) {
	c.mu.RLock()
	entry, success := c.Calls[key]
	c.mu.RUnlock()
	if success {
		return entry.val, success
	}
	return *new(T), success
}

func (c *Cache[T]) sweepLoop(retainSeconds int) {
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
