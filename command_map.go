package main

import (
	"fmt"
	"github.com/jake-abed/pokedexcli/internal/pokeapi"
)

func commandMap(config *commandConfig) error {
	if config.Next == nil && config.Previous != nil {
		fmt.Println("At end of map. Please try `mapb` command instead!")
		return nil
	}

	locationResp, err := config.pokeapiClient.ListLocations(config.Next)
	if err != nil {
		return err
	}

	locationAreas := locationResp.Results

	config.Next = locationResp.Next
	config.Previous = locationResp.Previous

	printLocationNames(locationAreas)
	return nil
}

func commandMapB(config *commandConfig) error {
	if config.Previous == nil {
		fmt.Println("At beginning of map. Please try `map` command instead!")
		return nil
	}

	locationResp, err := config.pokeapiClient.ListLocations(config.Previous)
	if err != nil {
		return err
	}

	locationAreas := locationResp.Results

	config.Next = locationResp.Next
	config.Previous = locationResp.Previous

	printLocationNames(locationAreas)
	return nil
}

func printLocationNames(locations []pokeapi.LocationAreas) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}
