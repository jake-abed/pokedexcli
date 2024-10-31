package pokeapi

import (
  "net/http"
  "encoding/json"
  "io"
  "errors"
)



func (c *Client) ListPokemonInLocation(location *string) ([]PokemonEncounters, error) {
  if location == nil {
    return nil, errors.New("A location must be provided.")
  }

  url := baseURL + "/" + *location

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }

  resp, err := c.httpClient.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  locationAreaData := RespLocationAreaData{}
  err = json.Unmarshal(data, &locationAreaData)
  if err != nil {
    return nil, err
  }

  return locationAreaData.PokemonEncounters, nil
}
