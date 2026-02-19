package main

import (
	"time"

	"github.com/Dr3iundZwanzig/pokedexcli/internal/pokecache"
)

func main() {
	cnf := config{
		Next:     "",
		Previous: "",
		Cache:    pokecache.NewCache(time.Minute * 5),
	}
	repl(&cnf)
}
