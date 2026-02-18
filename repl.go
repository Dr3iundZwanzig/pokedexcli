package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		userInput := scanner.Text()
		cleanUserInput := cleanInput(userInput)
		if len(cleanUserInput) == 0 {
			fmt.Print("Empty input\n")
			continue
		}
		fmt.Printf("Your command was: %v\n", cleanUserInput[0])
	}
}

func cleanInput(text string) []string {
	splitString := strings.Fields(strings.ToLower(text))
	return splitString
}
