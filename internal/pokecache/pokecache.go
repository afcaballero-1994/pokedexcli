package pokecache

import (
    "sync"
    "time"
)

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

type Cache struct {
    entry map[string]cacheEntry
    interval time.Duration
    mu *sync.Mutex
}

func NewCache(t time.Duration) Cache {
    c := Cache{
        entry: map[string]cacheEntry{},
        interval: t,
        mu: &sync.Mutex{},
    }
    go c.reapLoop()
    return c
}

func (c *Cache)Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.entry[key] = cacheEntry{createdAt: time.Now(), val: val,}
}

func (c *Cache)Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    data, exist := c.entry[key]
    return data.val, exist
}

func (c *Cache)reapLoop() {

    ticker := time.NewTicker(c.interval)
    defer ticker.Stop()

    c.mu.Lock()
    c.mu.Unlock()

    for {
        select {
            case t := <-ticker.C:
                for k, v := range c.entry{
                    diff := t.Sub(v.createdAt)
                    if diff > c.interval {
                        delete(c.entry, k)
                    }
                }
        }
    }
}
