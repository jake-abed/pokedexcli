package main

import (
  "os"
  "bufio"
  "fmt"
)

type cliCommand struct {
  name string
  description string
  callback func() error
}

func buildCommands() map[string]cliCommand {
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
  }
  return commands
}

func commandExit() error {
  fmt.Println("Exiting Pokedex... Goodbye for now!")
  return nil
}

func commandHelp() error {
  commands := buildCommands()
  for _, v := range commands {
    fmt.Println(fmt.Sprintf("%v - %v", v.name, v.description))
  }
  return nil
}

func main() {
  running := true
  commands := buildCommands()


  fmt.Println("Pokedex booting up!")
  fmt.Println("...")
  fmt.Println("Welcome to the Pokedex!")

  for running {
    
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Pokedex > ")

    scan := scanner.Scan()
    if scan {
      command, ok := commands[scanner.Text()]
      if ok {
        switch command.name {
        case "exit":
          command.callback()
          running = false
        default:
          command.callback()
        }
      } else {
        fmt.Println("Command not found! Please enter 'help' for aid.")
      }
    }
  }
}

