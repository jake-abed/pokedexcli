package main

import (
  "os"
  "bufio"
  "fmt"
  "strings"
)

type cliCommand struct {
  name string
  description string
  callback func(*commandConfig) error
}

type commandConfig struct {
  Next *string
  Previous *string
}

func buildCommands() (map[string]cliCommand, commandConfig) {
  config := commandConfig{}

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
          command.callback(&config)
          running = false
        default:
          command.callback(&config)
        }
      } else {
        fmt.Println("Command not found! Please enter 'help' for aid.")
      }
    }
  }
}

