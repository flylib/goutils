package cache

import (
	"container/list"
	"sync"
)

// LRUCache is an LRU cache. It is safe for concurrent access.
type LRUCache struct {
	mutex       sync.RWMutex
	maxItemSize int
	cacheList   *list.List
	cacheMap    map[interface{}]*list.Element
}

type pair struct {
	key   interface{}
	value interface{}
}

//NewLRUCache If maxItemSize is zero, the cache has no limit.
//if maxItemSize is not zero, when cache's size beyond maxItemSize,start to swap
func NewLRUCache(maxItemSize int) *LRUCache {
	return &LRUCache{
		maxItemSize: maxItemSize,
		cacheList:   list.New(),
		cacheMap:    make(map[interface{}]*list.Element),
	}
}

//Get value with key
func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if ele, hit := c.cacheMap[key]; hit {
		c.cacheList.MoveToFront(ele)
		return ele.Value.(*pair).value, true
	}
	return nil, false
}

//Set a value with key
func (c *LRUCache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.cacheMap == nil {
		c.cacheMap = make(map[interface{}]*list.Element)
		c.cacheList = list.New()
	}

	if ele, ok := c.cacheMap[key]; ok {
		c.cacheList.MoveToFront(ele)
		ele.Value.(*pair).value = value
		return
	}

	ele := c.cacheList.PushFront(&pair{key: key, value: value})
	c.cacheMap[key] = ele
	if c.maxItemSize != 0 && c.cacheList.Len() > c.maxItemSize {
		c.remOldest()
	}
}

//Delete delete the key
func (c *LRUCache) Del(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.cacheMap == nil {
		return
	}
	if ele, ok := c.cacheMap[key]; ok {
		c.cacheList.Remove(ele)
		key := ele.Value.(*pair).key
		delete(c.cacheMap, key)
		return
	}
}

//RemoveOldest remove the oldest key
func (c *LRUCache) remOldest() {
	if c.cacheMap == nil {
		return
	}
	ele := c.cacheList.Back()
	if ele != nil {
		c.cacheList.Remove(ele)
		key := ele.Value.(*pair).key
		delete(c.cacheMap, key)
	}
}
