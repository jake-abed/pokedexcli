package pokeapi

type RespPokeLocations struct {
  Count int `json:"count"`
  Next *string `json:"next"`
  Previous *string `json:"previous"`
  Results []LocationAreas `json:"results"`
}

type LocationAreas struct {
  Name string `json:"name"`
  URL string `json:"url"`
}

