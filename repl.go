package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Dr3iundZwanzig/pokedexcli/internal/pokecache"
)

func repl(cnf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		userInput := scanner.Text()
		cleanUserInput := cleanInput(userInput)

		commanName := cleanUserInput[0]
		secondInput := ""
		if len(userInput) == 0 {
			fmt.Print("Empty input\n")
			continue
		}
		if len(cleanUserInput) == 2 {
			secondInput = cleanUserInput[1]
		}
		if len(cleanUserInput) > 2 {
			fmt.Println("Too many inputs")
			continue
		}

		command, ok := getCommands()[commanName]
		if ok {
			err := command.callback(cnf, secondInput)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	splitString := strings.Fields(strings.ToLower(text))
	return splitString
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	Next     string
	Previous string
	Cache    pokecache.Cache
	Pokedex  map[string]Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Programm",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Opens the help page",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore the specified location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch the Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Try to catch the Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Try to catch the Pokemon",
			callback:    commandPokedex,
		},
	}
}
