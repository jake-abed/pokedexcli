package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespPokeLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

  cachedRes, ok := c.cache.Get(url)
  if ok {
    locations := RespPokeLocations{}
    err := json.Unmarshal(cachedRes, &locations)
    if err != nil {
      return RespPokeLocations{}, err
    }
    return locations, nil
  }

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokeLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokeLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokeLocations{}, err
	}

  c.cache.Add(url, dat)

	locationsResp := RespPokeLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespPokeLocations{}, err
	}

	return locationsResp, nil
}

