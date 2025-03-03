package cache

import (
  "time"
  "sync"
)

type cacheEntry struct {
  createdAt time.Time
  val []byte
}

type Cache struct {
  entries map[string]cacheEntry
  mu *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
  c := Cache{
    entries: make(map[string]cacheEntry),
    mu: &sync.Mutex{},
  }
  go c.reapLoop(interval)
  return c
}

func (c Cache) Add(key string, val []byte) {
  c.mu.Lock()
  defer c.mu.Unlock()

  c.entries[key] = cacheEntry{
    createdAt: time.Now(),
    val: val,
  }
}

func (c Cache) Get(key string) ([]byte, bool){
  cacheEntry, found := c.entries[key]
  if !found {
    return []byte{}, false
  }
  return cacheEntry.val, found
}

func (c *Cache) reapLoop(interval time.Duration) {
  ticker := time.NewTicker(interval)
  defer ticker.Stop()
  for {
    select {
      case <- ticker.C:
        c.mu.Lock()
        for key, entry := range c.entries {
          if time.Since(entry.createdAt) > interval {
            delete(c.entries, key)
          }
        }
        c.mu.Unlock()
    }
  }
}
