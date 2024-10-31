package pokeapi

import (
  "net/http"
  "time"
  "github.com/jake-abed/pokedexcli/internal/pokecache"
)

type Client struct {
  httpClient http.Client
  cache *pokecache.Cache
}

func NewClient(timeout time.Duration, cache *pokecache.Cache) Client {
  return Client{
    httpClient: http.Client{
      Timeout: timeout,
    },
    cache: cache,
  }
}

