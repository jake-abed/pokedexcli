package main

import (
  "fmt"
  "io"
  "encoding/json"
  "net/http"
)

type LocationAreaRes struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
  Results  []LocationResult `json:"results"`
}

type LocationResult struct {
		Name string `json:"name"`
		URL  string `json:"url"`
}

func commandMap(config *commandConfig) error {
  defaultUrl := "https://pokeapi.co/api/v2/location-area/"
  var targetUrl string

  if config.Next != nil {
    targetUrl = *config.Next
  } else if config.Previous == nil {
    targetUrl = defaultUrl
  } else {
    fmt.Println("End of the map reached. Please use `mapb` command.")
    return nil
  }

  res, err := http.Get(targetUrl)
  if err != nil {
    return err
  }
  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
    return err
  }

  var locationAreas LocationAreaRes
  if err := json.Unmarshal(body, &locationAreas); err != nil {
    return err
  }

  config.Next = locationAreas.Next
  config.Previous = locationAreas.Previous

  printLocationNames(locationAreas.Results)
  return nil
}

func commandMapB(config *commandConfig) error {
  var targetUrl string
  if config.Previous != nil {
    targetUrl = *config.Previous
  } else {
    fmt.Println("At beginning of map. Please try `map` command instead!")
    return nil
  }

  res, err := http.Get(targetUrl)
  if err != nil {
    return err
  }
  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
    return err
  }

  var locationAreas LocationAreaRes
  if err:= json.Unmarshal(body, &locationAreas); err != nil {
    return err
  }

  config.Next = locationAreas.Next
  config.Previous = locationAreas.Previous

  printLocationNames(locationAreas.Results)
  return nil
}

func printLocationNames (locations []LocationResult) {
  for _, location := range locations {
    fmt.Println(location.Name)
  }
}
