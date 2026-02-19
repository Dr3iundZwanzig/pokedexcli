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
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = cnf.Next
	}
	locations := getLocationAreas(url)
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	if locations.Previous != nil {
		cnf.Previous = *locations.Previous
	}
	if locations.Next != nil {
		cnf.Next = *locations.Next
	}

	return nil
}

func commandMapB(cnf *config) error {
	url := cnf.Previous
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locations := getLocationAreas(url)
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}

	if locations.Previous != nil {
		cnf.Previous = *locations.Previous
	} else {
		cnf.Previous = ""
	}
	if locations.Next != nil {
		cnf.Next = *locations.Next
	}
	return nil
}
