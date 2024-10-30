package main

import (
  "os"
  "fmt"
)

func commandExit(_ *commandConfig) error {
  fmt.Println("Exiting Pokedex... Goodbye for now!")
  os.Exit(0)
  return nil
}

