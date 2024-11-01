package pokeapi

import (
  "net/http"
  "encoding/json"
  "fmt"
  "io"
  "errors"
)



func (c *Client) ListPokemonInLocation(location *string) ([]PokemonEncounters, error) {
  if location == nil {
    return nil, errors.New("A location must be provided.")
  }

  url := baseURL + "/location-area/" + *location
  cachedRes, ok := c.cache.Get(url)
  if ok {
    encounters := RespLocationAreaData{}
    err := json.Unmarshal(cachedRes, &encounters)
    if err != nil {
      return nil, err
    }
    return encounters.PokemonEncounters, nil
  }

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }

  resp, err := c.httpClient.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  if resp.StatusCode == 404 {
    msg := fmt.Sprintf("%v is not a valid location you can explore.", *location)
    return nil, errors.New(msg)
  }

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  c.cache.Add(url, data)

  locationAreaData := RespLocationAreaData{}
  err = json.Unmarshal(data, &locationAreaData)
  if err != nil {
    return nil, err
  }

  return locationAreaData.PokemonEncounters, nil
}
