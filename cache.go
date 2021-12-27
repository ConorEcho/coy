package coy

import (
	"sync"

	"github.com/conorecho/coy/lru"
)

type cache struct {
	mu  sync.Mutex
	lru *lru.Cache

	cacheNum int
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lru == nil {
		c.lru = lru.New(c.cacheNum)
	}

	c.lru.Add(key, value)
}

func (c *cache) get(key string) (view ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}

	return
}
