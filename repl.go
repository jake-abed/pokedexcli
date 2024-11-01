package main

import (
  "os"
  "bufio"
  "fmt"
  "time"
  "strings"
  "github.com/jake-abed/pokedexcli/internal/pokeapi"
  "github.com/jake-abed/pokedexcli/internal/pokecache"
)

type cliCommand struct {
  name string
  description string
  callback func(*commandConfig) error
}

type commandConfig struct {
  pokeapiClient pokeapi.Client
  Next *string
  Previous *string
  Location *string
  Pokemon *string
}

func buildCommands() (map[string]cliCommand, *commandConfig) {
  cache := pokecache.NewCache(time.Second * 60)
  pokeClient := pokeapi.NewClient(5 * time.Second, &cache)
  config := &commandConfig{
    pokeapiClient: pokeClient,
  }

  commands := map[string]cliCommand{
    "help": {
      name: "help",
      description: "Displays a help message.",
      callback: commandHelp,
    },
    "exit": {
      name: "exit",
      description: "Exit the Pokedex.",
      callback: commandExit,
    },
    "map": {
      name: "map",
      description: "Displays the next 20 locations.",
      callback: commandMap,
    },
    "mapb": {
      name: "mapb",
      description: "Displays the previous 20 locations.",
      callback: commandMapB,
    },
    "explore": {
      name: "explore",
      description: "Explores a map location and possible pokemon encounters.",
      callback: commandExplore,
    },
    "catch": {
      name: "catch",
      description: "Attempt to catch a specific pokemon by name.",
      callback: commandCatch,
    },
  }
  return commands, config
}

func cleanInput(text string) []string {
  output := strings.ToLower(text)
  words := strings.Fields(output)
  return words
}

func runCli() {
  running := true
  commands, config := buildCommands()


  fmt.Println("Pokedex booting up...")
  fmt.Println("Success! Welcome to the Pokedex :)")

  for running {
    
  scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Pokedex > ")

    scan := scanner.Scan()
    if scan {
      input := cleanInput(scanner.Text())
      if len(input) == 0 {
        continue
      }
      command, ok := commands[input[0]]
      if ok {
        switch command.name {
        case "exit":
          err := command.callback(config)
          if err != nil {
            fmt.Println(err)
          }
          running = false
        case "explore":
          if len(input) < 2 {
            fmt.Println("Explore requires a location! Please try again.")
            continue
          }
          config.Location = &input[1]
          err := command.callback(config)
          if err != nil {
            fmt.Println(err)
          }
        case "catch":
          if len(input) < 2 {
            fmt.Println("Catch requires a pokemon name! Please try again.")
            continue
          }
          config.Pokemon = &input[1]
          err := command.callback(config)
          if err != nil {
            fmt.Println(err)
          }
        default:
          err := command.callback(config)
          if err != nil {
            fmt.Println(err)
          }
        }
      } else {
        fmt.Println("Command not found! Please enter 'help' for aid.")
      }
    }
  }
}

