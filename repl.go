package main

import (
	"strings"
)

func cleanInput(text string) []string {
	splitString := strings.Fields(strings.ToLower(text))
	return splitString
}
