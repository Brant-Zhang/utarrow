package ttlevent

import (
	"sync"
	"time"
)

//Cache contains item timing info for event creater
type Cache struct {
	pool  map[string]int64
	mu    sync.RWMutex
	ttl   int64
	event chan string
}

func newCache(ttl int) *Cache {
	m := new(Cache)
	m.pool = make(map[string]int64)
	m.ttl = int64(ttl)
	m.event = make(chan string, 1000)
	go m.check()
	return m
}

func (c *Cache) expire(t int64) bool {
	return time.Now().Before(time.Unix(t+c.ttl, 0))
}

//Put add itme in container
func (c *Cache) Put(key string) {
	c.mu.Lock()
	if _, ok := c.pool[key]; !ok {
		c.event <- "1/" + key
	}
	c.pool[key] = time.Now().Unix()
	c.mu.Unlock()
}

func (c *Cache) check() {
	tk := time.Tick(time.Duration(c.ttl * 1e9))
	for range tk {
		c.mu.Lock()
		for k, v := range c.pool {
			if c.expire(v) {
				c.event <- "0/" + k
				delete(c.pool, k)
			}
		}
		c.mu.Unlock()
	}
}
