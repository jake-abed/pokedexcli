package pokeapi

import (
  "net/http"
  "errors"
  "fmt"
  "io"
  "encoding/json"
)

func (c *Client) GetPokemonData(name *string) (PokemonData, error) {
  if name == nil {
    return PokemonData{}, errors.New("Please provide a Pokemon name")
  }

  url := baseURL + "/pokemon/" + *name

  cachedResp, ok := c.cache.Get(url)
  if ok {
    pokeData := PokemonData{}
    err := json.Unmarshal(cachedResp, &pokeData)
    if err != nil {
      return PokemonData{}, err
    }
    return pokeData, nil
  }

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return PokemonData{}, err
  }

  resp, err := c.httpClient.Do(req)
  if err != nil {
    return PokemonData{}, err
  }
  defer resp.Body.Close()
  
  if resp.StatusCode == 404 {
    msg := fmt.Sprintf("Hmmm... %v is not a valid pokemon... Try again!", *name)
    return PokemonData{}, errors.New(msg)
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return PokemonData{}, err
  }

  c.cache.Add(url, data)

  pokeData := PokemonData{}
  err = json.Unmarshal(data, &pokeData)
  if err != nil {
    return PokemonData{}, err
  }
  return pokeData, nil
}
