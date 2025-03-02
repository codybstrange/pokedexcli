package cache

import (
  "time"
)

type cacheEntry struct {
  createdAt time.Time
  val []byte
}

type Cache struct {
  entry map[string]cacheEntry
}

func NewCache() Cache {
  return Cache{}
}
