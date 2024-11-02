package main

import "fmt"

func commandHelp(_ *commandConfig) error {
	commands := buildCommands()
	fmt.Println("Welcome to this Pokedex!")
	fmt.Println("/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\")
	fmt.Println("How to use the Pokedex:")
	for _, v := range commands {
		fmt.Println("---")
		fmt.Println(fmt.Sprintf("%-8v <=>   %v", v.name, v.description))
	}
	fmt.Println("---")
	fmt.Println("If more help is required, please visit a PokeCenter!")
	return nil
}
