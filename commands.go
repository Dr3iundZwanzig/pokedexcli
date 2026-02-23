package main

import (
	"fmt"
	"math/rand/v2"
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
		return fmt.Errorf("Area not found: %v", si)
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

func commandCatch(cnf *config, si string) error {
	if si == "" {
		return fmt.Errorf("Second input empty")
	}

	pokemon, err := cnf.getPokemon(si)
	if err != nil {
		return fmt.Errorf("Pokemon not found: %v", si)
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	catchChance := 60 + pokemon.BaseExperience
	if rand.IntN(catchChance) > pokemon.BaseExperience {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		cnf.Pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}
	return nil
}

func commandInspect(cnf *config, si string) error {
	if si == "" {
		return fmt.Errorf("Second input empty")
	}

	pokemon, ok := cnf.Pokedex[si]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, Types := range pokemon.Types {
		fmt.Printf("  - %v\n", Types.Type.Name)
	}

	return nil
}

func commandPokedex(cnf *config, si string) error {
	if len(cnf.Pokedex) == 0 {
		return fmt.Errorf("Your Pokedex is empty")
	}

	fmt.Println("Your Pokedex:")
	for _, poke := range cnf.Pokedex {
		fmt.Printf(" - %v\n", poke.Name)
	}
	return nil
}
