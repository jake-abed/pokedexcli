package main

import (
  "fmt"
)

func commandExplore(c *commandConfig) error {
  pokemonEncounters, err := c.pokeapiClient.ListPokemonInLocation(c.Location)
  if err != nil {
    return err
  }

  fmt.Println(fmt.Sprintf("Exploring %v...", c.Location))
  for _, encounter := range pokemonEncounters {
    name := encounter.Pokemon.Name
    fmt.Println(fmt.Sprintf(" - %v", name))
  }
  return nil
}
