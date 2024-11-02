package main

import "fmt"

func commandPokedex(c *commandConfig) error {
  if len(c.pokeBank) == 0 {
    fmt.Println("There are no Pokemon in your Pokedex. Go catch 'em!")
    return nil
  }
  fmt.Println("The following Pokemon are in your Pokedex:")
  for k, _ := range c.pokeBank {
    fmt.Printf(" - %v\n", k)
  }
  return nil
}
