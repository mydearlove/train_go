package main

import (
	"container/list"
)

// LRUCache 是一个基于 LRU 算法的缓存结构
type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	lruList  *list.List
}

// CacheItem 存储在 LRUCache 中的数据项
type CacheItem struct {
	key   string
	value interface{}
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		lruList:  list.New(),
	}
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
	if elem, ok := c.cache[key]; ok {
		c.lruList.MoveToFront(elem)
		return elem.Value.(*CacheItem).value, true
	}
	return nil, false

}

func (c *LRUCache) Put(key string, value interface{}) {
	if elem, ok := c.cache[key]; ok {
		elem.Value.(*CacheItem).value = value
		c.lruList.MoveToFront(elem)
	}
	if len(c.cache) >= c.capacity {
		oldset := c.lruList.Back()
		if oldset != nil {
			delete(c.cache, oldset.Value.(*CacheItem).key)
			c.lruList.Remove(oldset)
		}
	}
	newElem := c.lruList.PushFront(&CacheItem{key: key, value: value})
	c.cache[key] = newElem

}
