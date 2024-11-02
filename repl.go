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
  pokeBank map[string]pokeapi.PokemonData
  Next *string
  Previous *string
  Location *string
  Pokemon *string
}

func buildConfig(cacheTime, clientTime time.Duration) (*commandConfig) {
  cache := pokecache.NewCache(time.Second * cacheTime)
  pokeClient := pokeapi.NewClient(time.Second * clientTime, &cache)
  return &commandConfig{
    pokeapiClient: pokeClient,
    pokeBank: make(map[string]pokeapi.PokemonData),
  }
}

func buildCommands() (map[string]cliCommand) {
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
      description: "Explores a map location and possible Pokemon encounters.",
      callback: commandExplore,
    },
    "catch": {
      name: "catch",
      description: "Attempt to catch a specific Pokemon by name.",
      callback: commandCatch,
    },
    "inspect": {
      name: "inspect",
      description: "Inspect a Pokemon in your PokeBank.",
      callback: commandInspect,
    },
    "pokedex": {
      name: "pokedex",
      description: "View all Pokemon caught so far.",
      callback: commandPokedex,
    },
  }
  return commands
}

func cleanInput(text string) []string {
  output := strings.ToLower(text)
  words := strings.Fields(output)
  return words
}

func runCli() {
  running := true
  commands := buildCommands()
  config := buildConfig(300, 5)

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
        case "explore":
          if len(input) < 2 {
            fmt.Println("Explore requires a location! Please try again.")
            continue
          }
          config.Location = &input[1]
          command.execute(config)
        case "catch":
          if len(input) < 2 {
            fmt.Println("Catch requires a pokemon name! Please try again.")
            continue
          }
          config.Pokemon = &input[1]
          command.execute(config)
        case "inspect":
          if len(input) < 2 {
            fmt.Println("Inspect requires a pokemon name! Please try again.")
            continue
          }
          config.Pokemon = &input[1]
          command.execute(config)
        default:
          command.execute(config)
        }
      } else {
        fmt.Println("Command not found! Please enter 'help' for aid.")
      }
    }
  }
}

func (c *cliCommand) execute(cfg *commandConfig) {
  err := c.callback(cfg)
  if err != nil {
    fmt.Println(err)
  }
}

