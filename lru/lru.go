package lru

import "container/list"

type Cache struct {
	ll    *list.List
	cache map[string]*list.Element

	curSize int
	maxSize int
}

type entry struct {
	k string
	v interface{}
}

func New(size int) *Cache {
	return &Cache{
		ll:      list.New(),
		cache:   make(map[string]*list.Element),
		curSize: 0,
		maxSize: size,
	}
}

func (c *Cache) remove(key string) {
	c.ll.Remove(c.cache[key])
	delete(c.cache, key)
	c.curSize--
}

func (c *Cache) getOldestEntry() entry {
	oldest := c.ll.Back()
	return oldest.Value.(entry)
}

func (c *Cache) add(key string, value interface{}) {
	elem := c.ll.PushFront(entry{k: key, v: value})
	c.cache[key] = elem
	c.curSize++
}

func (c *Cache) Add(key string, value interface{}) {
	if _, ok := c.Get(key); ok {
		c.remove(key)
	}

	if c.curSize >= c.maxSize {
		oldestEntry := c.getOldestEntry()
		c.remove(oldestEntry.k)
	}

	c.add(key, value)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		val := elem.Value.(entry)
		return val.v, true
	}

	return nil, false
}

func (c *Cache) Used() int {
	return c.curSize
}
