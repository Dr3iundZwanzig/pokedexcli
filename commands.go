package main

import (
	"fmt"
	"os"
)

func commandExit(cnf *config, si string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cnf *config, si string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func commandMap(cnf *config, si string) error {
	url := ""
	if cnf.Next == "end" {
		return fmt.Errorf("Last entry reached")
	}
	if cnf.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	} else {
		url = cnf.Next
	}
	locations, err := cnf.getLocationAreas(url)
	if err != nil {
		return err
	}
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapB(cnf *config, si string) error {
	url := cnf.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locations, err := cnf.getLocationAreas(url)
	if err != nil {
		return err
	}
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandExplore(cnf *config, si string) error {
	if si == "" {
		return fmt.Errorf("Second input empty")
	}
	locationArea, err := cnf.getLocationArea(si)
	if err != nil {
		return fmt.Errorf("Area not found")
	}
	fmt.Printf("Exploring %v\n", locationArea.Name)
	if len(locationArea.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon in that area")
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}
	return nil
}
