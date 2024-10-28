package main

import (
  "os"
  "bufio"
  "fmt"
)

func main() {
  running := true

  fmt.Println("Pokedex booting up!")
  fmt.Println("...")
  fmt.Println("Welcome to the Pokedex!")

  for running {
    
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Pokedex > ")

    scan := scanner.Scan()
    if scan {
      fmt.Println(scanner.Text())
    }
  }
}
