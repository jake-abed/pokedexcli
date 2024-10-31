package pokecache

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

