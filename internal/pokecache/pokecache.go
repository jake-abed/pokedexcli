package pokecache

import (
  "time"
  "sync"
)

func NewCache(interval time.Duration) Cache {
  cache := Cache{
    entries: make(map[string]cacheEntry),
    mu: &sync.Mutex{},
  }

  cache.reapLoop(interval)

  return cache
}

func (c *Cache) Add(key string, val []byte) {
  c.mu.Lock()
  defer c.mu.Unlock()

  entry := cacheEntry{
    createdAt: time.Now(),
    val: val,
  }

  c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
  c.mu.Lock()
  defer c.mu.Unlock()

  entry, ok := c.entries[key]
  if !ok {
    return []byte{}, false
  }
  return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
  ticker := time.NewTicker(interval)
  defer ticker.Stop()
  go func() {
    for range ticker.C {
      c.reap(time.Now(), interval)
    }
  }()
}

func (c *Cache) reap(now time.Time, last time.Duration) {
  c.mu.Lock()
  defer c.mu.Unlock()
  for k, v := range c.entries {
    if v.createdAt.Before(now.Add(-last)) {
      delete(c.entries, k)
    }
  }
}
