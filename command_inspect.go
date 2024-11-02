package main

import "fmt"

func commandInspect(c *commandConfig) error {
	pokemon, ok := c.pokeBank[*c.Pokemon]
	if !ok {
		fmt.Println("That pokemon is not in your PokeBank!")
		return nil
	}

	height := pokemon.Height
	weight := pokemon.Weight
	hp := pokemon.Stats[0].BaseStat
	attack := pokemon.Stats[1].BaseStat
	defense := pokemon.Stats[2].BaseStat
	specialAttack := pokemon.Stats[3].BaseStat
	specialDefense := pokemon.Stats[4].BaseStat
	speed := pokemon.Stats[5].BaseStat

	fmt.Printf("Name: %v\n", *c.Pokemon)
	fmt.Printf("Height: %v\n", height)
	fmt.Printf("Weight: %v\n", weight)
	fmt.Println("Stats:")
	fmt.Printf(" - HP: %d\n", hp)
	fmt.Printf(" - Attack: %d\n", attack)
	fmt.Printf(" - Defense: %d\n", defense)
	fmt.Printf(" - Special Attack: %d\n", specialAttack)
	fmt.Printf(" - Special Defense: %d\n", specialDefense)
	fmt.Printf(" - Speed: %d\n", speed)
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %v\n", t.Type.Name)
	}
	return nil
}
