package memory

import (
	"github.com/hashicorp/golang-lru/simplelru"
	"github.com/cnych/starjazz/cache"
	"github.com/cnych/starjazz/syncx"
	"sync"
	"time"
)

type Cache struct {
	l  *simplelru.LRU
	mu sync.RWMutex
}

func New(size int) (*Cache, error) {
	l, err := simplelru.NewLRU(size, nil)
	if err != nil {
		return nil, err
	}
	return &Cache{l, sync.RWMutex{}}, nil
}

// Keys
func (c *Cache) Keys() []string {
	var keys []string
	syncx.RLockDo(&c.mu, func() {
		for _, k := range c.l.Keys() {
			keys = append(keys, k.(string))
		}
	})
	return keys
}

// Get / MGet
func (c *Cache) Get(k string) (interface{}, error) {
	v0 := c.get(k)
	return getValue(v0, time.Now()), nil
}

func (c *Cache) MGet(ks []string) (map[string]interface{}, error) {
	n := len(ks)
	if n == 0 {
		return make(map[string]interface{}, 0), nil
	}
	v0m := c.mGet(ks)
	now := time.Now()
	vm := make(map[string]interface{}, n)
	for k, v0 := range v0m {
		vm[k] = getValue(v0, now)
	}
	return vm, nil
}

// Put / MPut
func (c *Cache) Put(k string, v interface{}, ttl time.Duration) error {
	if v == cache.TheMissing {
		return ErrIllegalValue
	}
	if v == nil {
		return nil
	}
	c.put(k, makeValue(v, ttl, time.Now()))
	return nil
}

func (c *Cache) MPut(kvs map[string]interface{}, ttl time.Duration) error {
	n := len(kvs)
	if n == 0 {
		return nil
	}

	now := time.Now()
	v0m := make(map[string]*value, n)
	for k, v := range kvs {
		if v == cache.TheMissing {
			return ErrIllegalValue
		}
		if v == nil {
			continue
		}
		v0m[k] = makeValue(v, ttl, now)
	}
	c.mPut(v0m)
	return nil
}

// MarkMissing / MMarkMissing
func (c *Cache) MarkMissing(k string, ttl time.Duration) error {
	c.put(k, makeValue(cache.TheMissing, ttl, time.Now()))
	return nil
}

func (c *Cache) MMarkMissing(ks []string, ttl time.Duration) error {
	n := len(ks)
	if n == 0 {
		return nil
	}

	now := time.Now()
	v0m := make(map[string]*value, n)
	for _, k := range ks {
		v0m[k] = makeValue(cache.TheMissing, ttl, now)
	}
	c.mPut(v0m)
	return nil
}

// Helpers
func (c *Cache) Remove(k string) error {
	syncx.WLockDo(&c.mu, func() {
		c.l.Remove(k)
	})
	return nil
}

func (c *Cache) MRemove(ks []string) error {
	if len(ks) == 0 {
		return nil
	}
	syncx.WLockDo(&c.mu, func() {
		for _, k := range ks {
			c.l.Remove(k)
		}
	})
	return nil
}

func (c *Cache) get(k string) *value {
	var v0 *value
	var ok bool
	syncx.WLockDo(&c.mu, func() { // 在LRU内部,get操作其实也要写内部数据，所以用WLock
		var v00 interface{}
		v00, ok = c.l.Get(k)
		if v00 != nil {
			v0 = v00.(*value)
		}
	})
	if ok {
		return v0
	} else {
		return nil
	}
}

func (c *Cache) mGet(ks []string) map[string]*value {
	v0m := make(map[string]*value, len(ks))
	syncx.WLockDo(&c.mu, func() { // 在LRU内部,get操作其实也要写内部数据，所以用WLock
		var v0 interface{}
		var ok bool
		for _, k := range ks {
			v0, ok = c.l.Get(k)
			if ok {
				v0m[k] = v0.(*value)
			} else {
				v0m[k] = nil
			}
		}
	})
	return v0m
}

func (c *Cache) put(k string, v0 *value) {
	syncx.WLockDo(&c.mu, func() {
		c.l.Add(k, v0)
	})
}

func (c *Cache) mPut(v0m map[string]*value) {
	syncx.WLockDo(&c.mu, func() {
		for k, v0 := range v0m {
			c.l.Add(k, v0)
		}
	})
}
