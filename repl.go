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
      description: "Displays a help message",
      callback: commandHelp,
    },
    "exit": {
      name: "exit",
      description: "Exit the Pokedex",
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
  }
  return commands, config
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
      command, ok := commands[strings.ToLower(scanner.Text())]
      if ok {
        switch command.name {
        case "exit":
          err := command.callback(config)
          if err != nil {
            fmt.Println(err)
          }
          running = false
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

