package httpcache

import (
	lru "github.com/hashicorp/golang-lru"
)

type LruCache struct {
	cache *lru.Cache
}

func NewLruCache(size int) *LruCache {
	c, _ := lru.New(size)
	return &LruCache{cache: c}
}

func (c *LruCache) Get(key string) ([]byte, bool) {
	v, ok := c.cache.Get(key)
	if !ok {
		return nil, false
	}
	return v.([]byte), true
}

func (c *LruCache) Set(key string, resp []byte) {
	c.cache.Add(key, resp)
}

func (c *LruCache) Delete(key string) {
	c.cache.Remove(key)
}
