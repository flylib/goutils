package cache

import (
	"fmt"
	"sync"
	"time"
)

// 并发安全内存库
type Cache[T any] struct {
	store sync.Map
}

type Item[T any] struct {
	expiration int64
	value      T
}

func NewCache[T any](clearInterval time.Duration) *Cache[T] {
	c := &Cache[T]{}
	if clearInterval > 0 {
		go func() {
			ticker := time.NewTicker(clearInterval)
			for range ticker.C {
				now := time.Now().UnixNano()
				c.store.Range(func(k, value any) bool {
					v := value.(*Item[T])
					if v.expiration < now {
						c.store.Delete(k)
					}
					return true
				})
				fmt.Println("clear done")
			}
		}()
	}
	return c
}

// 替换新值不覆盖过期时间
func (c *Cache[T]) SwapValue(k string, v T) bool {
	value, ok := c.store.Load(k)
	if !ok {
		return ok
	}
	item := value.(*Item[T])
	item.value = v
	return ok
}

// 存储，如果存在旧值则会覆盖过期时间
func (c *Cache[T]) Story(k string, v T, d time.Duration) {
	item := Item[T]{
		expiration: time.Now().Add(d).UnixNano(),
		value:      v,
	}
	c.store.Store(k, &item)
}

func (c *Cache[T]) Get(k string) (v T, ok bool) {
	value, ok := c.store.Load(k)
	if !ok {
		return
	}
	item := value.(*Item[T])
	if item.expiration < time.Now().UnixNano() {
		return v, false
	}
	return item.value, ok
}

func (c *Cache[T]) Delete(k string) {
	c.store.Delete(k)
}
