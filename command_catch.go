package main

import (
  "math/rand"
  "fmt"
  "time"
  "strings"
)

func commandCatch(c *commandConfig) error {
  pokeData, err := c.pokeapiClient.GetPokemonData(c.Pokemon)
  if err != nil {
    return err
  }

  baseXp := pokeData.BaseExperience
  name:= strings.ToUpper(pokeData.Name[:1]) + pokeData.Name[1:]
  throw := rand.Intn(400)
  
  // This section is to catch outlier pokemon who have more than
  // 340 baseXp and constrain them to that amount.
  if baseXp > 340 {
    baseXp = 340
  }

  fmt.Println(fmt.Sprintf("Throwing a pokeball at %v...", name))
  time.Sleep(500 * time.Millisecond)
  fmt.Print(".")
  time.Sleep(500 * time.Millisecond)
  fmt.Print(".")
  time.Sleep(500 * time.Millisecond)
  fmt.Print(".\n")
  if throw >= baseXp {
    c.pokeBank[pokeData.Name] = pokeData
    msg := fmt.Sprintf("Awesome! You caught %v!", name)
    fmt.Println(msg)
    return nil
  } else {
    msg := fmt.Sprintf("Dang! %v got away...", name)
    fmt.Println(msg)
    return nil
  }
}

