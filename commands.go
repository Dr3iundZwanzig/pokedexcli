package main

import (
	"fmt"
	"os"
)

func commandExit(cnf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cnf *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}
	return nil
}

func commandMap(cnf *config) error {
	url := ""
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

func commandMapB(cnf *config) error {
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
